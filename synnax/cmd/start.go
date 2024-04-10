// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package cmd

import (
	"bufio"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/synnaxlabs/synnax/pkg/hardware"
	"github.com/synnaxlabs/synnax/pkg/label"
	"github.com/synnaxlabs/synnax/pkg/ranger"
	"github.com/synnaxlabs/synnax/pkg/version"
	"github.com/synnaxlabs/synnax/pkg/workspace"
	"github.com/synnaxlabs/synnax/pkg/workspace/vis"

	"github.com/samber/lo"
	"github.com/synnaxlabs/synnax/pkg/security"
	"google.golang.org/grpc/credentials"
	insecureGRPC "google.golang.org/grpc/credentials/insecure"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/synnaxlabs/alamos"
	"github.com/synnaxlabs/freighter/fgrpc"
	"github.com/synnaxlabs/freighter/fhttp"
	"github.com/synnaxlabs/synnax/pkg/access"
	"github.com/synnaxlabs/synnax/pkg/api"
	grpcapi "github.com/synnaxlabs/synnax/pkg/api/grpc"
	httpapi "github.com/synnaxlabs/synnax/pkg/api/http"
	"github.com/synnaxlabs/synnax/pkg/auth"
	"github.com/synnaxlabs/synnax/pkg/auth/password"
	"github.com/synnaxlabs/synnax/pkg/auth/token"
	"github.com/synnaxlabs/synnax/pkg/distribution"
	"github.com/synnaxlabs/synnax/pkg/server"
	"github.com/synnaxlabs/synnax/pkg/storage"
	"github.com/synnaxlabs/synnax/pkg/user"
	"github.com/synnaxlabs/x/address"
	"github.com/synnaxlabs/x/config"
	"github.com/synnaxlabs/x/gorp"
	xsignal "github.com/synnaxlabs/x/signal"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a Synnax Node",
	Long: `
Starts a Synnax Node using the data directory specified by the --data flag,
and listening on the address specified by the --listen flag. If --peers
is specified and no existing data is found, the node will attempt to join the cluster
formed by its peers. If no peers are specified and no existing data is found, the node
will bootstrap a new cluster.
	`,
	Example: `synnax start --listen [host:port] --data /mnt/ssd1 --peers [host:port],[host:port] --insecure`,
	Args:    cobra.NoArgs,
	Run:     func(cmd *cobra.Command, _ []string) { start(cmd) },
}

var (
	stopKeyWord = "stop"
)

// start a Synnax node using the configuration specified by the command line flags,
// environment variables, and configuration files.
func start(cmd *cobra.Command) {
	v := version.Get()
	var (
		ins      = configureInstrumentation(v)
		insecure = viper.GetBool("insecure")
		verbose  = viper.GetBool("verbose")
		autoCert = viper.GetBool("auto-cert")
	)
	defer cleanupInstrumentation(cmd.Context(), ins)

	if autoCert {
		if err := generateAutoCerts(ins); err != nil {
			ins.L.Fatal("failed to generate auto certs", zap.Error(err))
		}
	}

	ins.L.Info("starting Synnax node", zap.String("version", v))

	interruptC := make(chan os.Signal, 1)
	signal.Notify(interruptC, os.Interrupt)

	// Any data stored on the node is considered sensitive, so we need to set the
	// permission mask for all files appropriately.
	disablePermissionBits()

	sCtx, cancel := xsignal.WithCancel(cmd.Context(), xsignal.WithInstrumentation(ins))
	defer cancel()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == stopKeyWord {
				interruptC <- os.Interrupt
			}
		}
	}()

	// Perform the rest of the startup within a separate goroutine, so we can properly
	// handle signal interrupts.
	sCtx.Go(func(ctx context.Context) error {

		secProvider, err := configureSecurity(ins, insecure)
		if err != nil {
			return err
		}

		// An array to hold the grpcTransports we use for cluster internal communication.
		grpcTransports := &[]fgrpc.BindableTransport{}

		grpcPool := configureClientGRPC(secProvider, insecure)

		// Open the distribution layer.
		storageCfg := buildStorageConfig(ins)
		distConfig, err := buildDistributionConfig(
			grpcPool,
			ins,
			storageCfg,
			grpcTransports,
		)
		dist, err := distribution.Open(ctx, distConfig)
		if err != nil {
			return err
		}
		defer func() { err = dist.Close() }()

		// set up our high level services.
		gorpDB := dist.Storage.Gorpify()
		userSvc, err := user.NewService(ctx, user.Config{
			DB:       gorpDB,
			Ontology: dist.Ontology,
			Group:    dist.Group,
		})
		tokenSvc := &token.Service{KeyProvider: secProvider, Expiration: 24 * time.Hour}
		authenticator := &auth.KV{DB: gorpDB}
		rangeSvc, err := ranger.OpenService(ctx, ranger.Config{
			DB:       gorpDB,
			Ontology: dist.Ontology,
			Group:    dist.Group,
			Signals:  dist.Signals,
		})
		if err != nil {
			return err
		}
		workspaceSvc, err := workspace.NewService(ctx, workspace.Config{DB: gorpDB, Ontology: dist.Ontology, Group: dist.Group})
		if err != nil {
			return err
		}
		visSvc, err := vis.NewService(vis.Config{DB: gorpDB, Ontology: dist.Ontology})
		if err != nil {
			return err
		}
		labelSvc, err := label.OpenService(ctx, label.Config{
			DB:       gorpDB,
			Ontology: dist.Ontology,
			Group:    dist.Group,
			Signals:  dist.Signals,
		})
		deviceSvc, err := hardware.OpenService(ctx, hardware.Config{
			DB:           gorpDB,
			Ontology:     dist.Ontology,
			Group:        dist.Group,
			HostProvider: dist.Cluster,
			Signals:      dist.Signals,
			Channel:      dist.Channel,
		})
		if err != nil {
			return err
		}

		// Provision the root user.
		if err := maybeProvisionRootUser(ctx, gorpDB, authenticator, userSvc); err != nil {
			return err
		}

		// Configure the API core.
		_api, err := api.New(api.Config{
			Instrumentation: ins.Child("api"),
			Authenticator:   authenticator,
			Enforcer:        access.AllowAll{},
			Vis:             visSvc,
			Insecure:        config.Bool(insecure),
			Channel:         dist.Channel,
			Framer:          dist.Framer,
			Storage:         dist.Storage,
			User:            userSvc,
			Token:           tokenSvc,
			Cluster:         dist.Cluster,
			Ontology:        dist.Ontology,
			Group:           dist.Group,
			Ranger:          rangeSvc,
			Workspace:       workspaceSvc,
			Label:           labelSvc,
			Hardware:        deviceSvc,
		})
		if err != nil {
			return err
		}

		// Configure the HTTP API Transport.
		r := fhttp.NewRouter(fhttp.RouterConfig{Instrumentation: ins})
		_api.BindTo(httpapi.New(r))

		// Configure the GRPC API Transport.
		grpcAPI, grpcAPITrans := grpcapi.New()
		*grpcTransports = append(*grpcTransports, grpcAPITrans...)
		_api.BindTo(grpcAPI)

		srv, err := server.New(buildServerConfig(
			*grpcTransports,
			[]fhttp.BindableTransport{r},
			secProvider,
			ins,
			verbose,
		))
		if err != nil {
			return err
		}
		sCtx.Go(func(_ context.Context) error {
			defer cancel()
			return srv.Serve()
		}, xsignal.WithKey("server"))
		defer srv.Stop()
		<-ctx.Done()
		return nil
	}, xsignal.WithKey("start"))

	select {
	case <-interruptC:
		ins.L.Info("received interrupt signal, shutting down")
		cancel()
	case <-sCtx.Stopped():
	}

	if err := sCtx.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		ins.L.Fatal("synnax failed", zap.Error(err))
	}
	ins.L.Info("shutdown successful")
}

func init() {
	rootCmd.AddCommand(startCmd)
	configureStartFlags()
	bindFlags(startCmd)
}

func buildStorageConfig(
	ins alamos.Instrumentation,
) storage.Config {
	return storage.Config{
		Instrumentation: ins.Child("storage"),
		MemBacked:       config.Bool(viper.GetBool("mem")),
		Dirname:         viper.GetString("data"),
	}
}

func parsePeerAddresses() ([]address.Address, error) {
	peerStrings := viper.GetStringSlice("peers")
	peerAddresses := make([]address.Address, len(peerStrings))
	for i, listenString := range peerStrings {
		peerAddresses[i] = address.Address(listenString)
	}
	return peerAddresses, nil
}

func buildDistributionConfig(
	pool *fgrpc.Pool,
	ins alamos.Instrumentation,
	storage storage.Config,
	transports *[]fgrpc.BindableTransport,
) (distribution.Config, error) {
	peers, err := parsePeerAddresses()
	return distribution.Config{
		Instrumentation:  ins.Child("distribution"),
		AdvertiseAddress: address.Address(viper.GetString("listen")),
		PeerAddresses:    peers,
		Pool:             pool,
		Storage:          storage,
		Transports:       transports,
	}, err
}

func buildServerConfig(
	grpcTransports []fgrpc.BindableTransport,
	httpTransports []fhttp.BindableTransport,
	sec security.Provider,
	ins alamos.Instrumentation,
	debug bool,
) (cfg server.Config) {
	cfg.Branches = append(cfg.Branches,
		&server.SecureHTTPBranch{Transports: httpTransports},
		&server.GRPCBranch{Transports: grpcTransports},
		server.NewHTTPRedirectBranch(),
	)
	cfg.Debug = config.Bool(debug)
	cfg.ListenAddress = address.Address(viper.GetString("listen"))
	cfg.Instrumentation = ins.Child("server")
	cfg.Security.TLS = sec.TLS()
	cfg.Security.Insecure = config.Bool(viper.GetBool("insecure"))
	return cfg
}

func configureSecurity(ins alamos.Instrumentation, insecure bool) (security.Provider, error) {
	return security.NewProvider(security.ProviderConfig{
		LoaderConfig: buildCertLoaderConfig(ins),
		Insecure:     config.Bool(insecure),
		KeySize:      viper.GetInt("key-size"),
	})
}

func maybeProvisionRootUser(
	ctx context.Context,
	db *gorp.DB,
	authSvc auth.Authenticator,
	userSvc *user.Service,
) error {
	creds := auth.InsecureCredentials{
		Username: viper.GetString("username"),
		Password: password.Raw(viper.GetString("password")),
	}
	exists, err := userSvc.UsernameExists(ctx, creds.Username)
	if err != nil || exists {
		return err
	}
	return db.WithTx(ctx, func(tx gorp.Tx) error {
		if err = authSvc.NewWriter(tx).Register(ctx, creds); err != nil {
			return err
		}
		return userSvc.NewWriter(tx).Create(ctx, &user.User{Username: creds.Username})
	})
}

func configureClientGRPC(
	sec security.Provider,
	insecure bool,
) *fgrpc.Pool {
	return fgrpc.NewPool(
		grpc.WithTransportCredentials(getClientGRPCTransportCredentials(sec, insecure)),
	)
}

func getClientGRPCTransportCredentials(sec security.Provider, insecure bool) credentials.TransportCredentials {
	return lo.Ternary(insecure, insecureGRPC.NewCredentials(), credentials.NewTLS(sec.TLS()))
}
