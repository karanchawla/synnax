// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package cluster

import (
	"github.com/synnaxlabs/alamos"
	"github.com/synnaxlabs/aspen/internal/cluster/gossip"
	pledge_ "github.com/synnaxlabs/aspen/internal/cluster/pledge"
	"github.com/synnaxlabs/x/address"
	"github.com/synnaxlabs/x/binary"
	"github.com/synnaxlabs/x/config"
	"github.com/synnaxlabs/x/kv"
	"github.com/synnaxlabs/x/override"
	"github.com/synnaxlabs/x/validate"
	"time"
)

const FlushOnEvery = -1 * time.Second

type Config struct {
	alamos.Instrumentation
	// HostAddress is the reachable address of the host node.
	// [REQUIRED]
	HostAddress address.Address
	// Storage is a key-value storage backend for the cluster. Cluster will flush
	// changes to its state to this backend based on Config.StorageFlushInterval.
	// Open will also attempt to load an existing cluster from this backend.
	// If Config.Storage is not provided, Cluster state will only be stored in memory.
	Storage kv.DB
	// StorageKey is the key used to store the cluster state in the backend.
	StorageKey []byte
	// StorageFlushInterval	is the interval at which the cluster state is flushed
	// to the backend. If this is set to FlushOnEvery, the cluster state is flushed on
	// every change.
	StorageFlushInterval time.Duration
	// Gossip is the configuration for propagating Cluster state through gossip.
	// See the gossip package for more details on how to configure this.
	Gossip gossip.Config
	// Pledge is the configuration for pledging to the cluster upon a Open call.
	// See the pledge package for more details on how to configure this.
	Pledge pledge_.Config
	// EncoderDecoder is the encoder/decoder to use for encoding and decoding the
	// cluster state.
	EncoderDecoder binary.EncoderDecoder
}

var _ config.Config[Config] = Config{}

func (cfg Config) Override(other Config) Config {
	cfg.HostAddress = override.String(cfg.HostAddress, other.HostAddress)
	cfg.EncoderDecoder = override.Nil(cfg.EncoderDecoder, other.EncoderDecoder)
	cfg.StorageFlushInterval = override.Numeric(cfg.StorageFlushInterval, other.StorageFlushInterval)
	cfg.StorageKey = override.Slice(cfg.StorageKey, other.StorageKey)
	cfg.Storage = override.Nil(cfg.Storage, other.Storage)
	cfg.Instrumentation = override.Zero(cfg.Instrumentation, other.Instrumentation)
	cfg.Gossip.Instrumentation = cfg.Instrumentation
	cfg.Pledge.Instrumentation = cfg.Instrumentation
	cfg.Gossip = cfg.Gossip.Override(other.Gossip)
	cfg.Pledge = cfg.Pledge.Override(other.Pledge)
	return cfg
}

func (cfg Config) Validate() error {
	v := validate.New("cluster")
	validate.NotEmptyString(v, "HostAddress", cfg.HostAddress)
	validate.NotNil(v, "EncoderDecoder", cfg.EncoderDecoder)
	validate.NonZero(v, "StorageFlushInterval", cfg.StorageFlushInterval)
	validate.NotEmptySlice(v, "StorageKey", cfg.StorageKey)
	return v.Error()
}

// Report implements the alamos.ReportProvider interface.
func (cfg Config) Report() alamos.Report {
	report := make(alamos.Report)
	if cfg.Storage != nil {
		report["storage"] = cfg.Storage.Report()
	} else {
		report["storage"] = "not provided"
	}
	report["storageKey"] = string(cfg.StorageKey)
	report["storageFlushInterval"] = cfg.StorageFlushInterval
	return report
}

var (
	DefaultConfig = Config{
		Pledge:               pledge_.DefaultConfig,
		StorageKey:           []byte("aspen.cluster"),
		Gossip:               gossip.DefaultConfig,
		StorageFlushInterval: 1 * time.Second,
		EncoderDecoder:       &binary.GobEncoderDecoder{},
	}
	FastConfig = DefaultConfig.Override(Config{
		Pledge: pledge_.FastConfig,
		Gossip: gossip.FastConfig,
	})
	BlazingFastConfig = DefaultConfig.Override(Config{
		Pledge: pledge_.BlazingFastConfig,
		Gossip: gossip.FastConfig,
	})
)
