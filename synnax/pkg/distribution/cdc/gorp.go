// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package cdc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/synnaxlabs/synnax/pkg/distribution/channel"
	"github.com/synnaxlabs/x/binary"
	"github.com/synnaxlabs/x/change"
	"github.com/synnaxlabs/x/config"
	"github.com/synnaxlabs/x/gorp"
	"github.com/synnaxlabs/x/observe"
	"github.com/synnaxlabs/x/override"
	"github.com/synnaxlabs/x/telem"
	"github.com/synnaxlabs/x/types"
	"github.com/synnaxlabs/x/validate"
	"go.uber.org/zap"
	"io"
	"strings"
)

// GorpConfig is the configuration for opening a CDC pipeline that subscribes
// changes to a particular entry type in a gorp.DB. It's not typically necessary
// to instantiate this configuration directly, instead use a helper function
// such as GorpConfigUUID.
type GorpConfig[K gorp.Key, E gorp.Entry[K]] struct {
	// DB is the DB to subscribe to.
	DB *gorp.DB
	// SetDataType is the data type of the key used by the DB.
	SetDataType telem.DataType
	// DeleteDataType is the data type of the key used by the DB.
	DeleteDataType telem.DataType
	// MarshalSet is a function that marshals the key used by the DB into a byte slice.
	MarshalSet func(entry E) ([]byte, error)
	// MarshalDelete is a function that marshals the key used by the DB into a byte slice.
	MarshalDelete func(K) ([]byte, error)
	// SetName is the name of the set channel.
	SetName string
	// DeleteName is the name of the delete channel.
	DeleteName string
}

var _ config.Config[GorpConfig[uuid.UUID, gorp.Entry[uuid.UUID]]] = GorpConfig[uuid.UUID, gorp.Entry[uuid.UUID]]{}

func DefaultGorpConfig[K gorp.Key, E gorp.Entry[K]]() GorpConfig[K, E] {
	t := types.Name[E]()
	return GorpConfig[K, E]{
		SetName:    fmt.Sprintf("sy_%s_set", strings.ToLower(t)),
		DeleteName: fmt.Sprintf("sy_%s_delete", strings.ToLower(t)),
	}
}

func (g GorpConfig[K, E]) Override(other GorpConfig[K, E]) GorpConfig[K, E] {
	g.DB = override.Nil(g.DB, other.DB)
	g.SetDataType = override.String(g.SetDataType, other.SetDataType)
	g.DeleteDataType = override.String(g.DeleteDataType, other.DeleteDataType)
	g.MarshalSet = override.Nil(g.MarshalSet, other.MarshalSet)
	g.MarshalDelete = override.Nil(g.MarshalDelete, other.MarshalDelete)
	g.SetName = override.String(g.SetName, other.SetName)
	g.DeleteName = override.String(g.DeleteName, other.DeleteName)
	return g
}

func (g GorpConfig[K, E]) Validate() error {
	v := validate.New("cdc.GorpConfig")
	validate.NotEmptyString(v, "SetName", g.SetName)
	validate.NotEmptyString(v, "DeleteName", g.DeleteName)
	validate.NotNil(v, "DB", g.DB)
	validate.NotEmptyString(v, "SetDataType", g.SetDataType)
	validate.NotEmptyString(v, "DeleteDataType", g.DeleteDataType)
	validate.NotNil(v, "MarshalSet", g.MarshalSet)
	validate.NotNil(v, "MarshalDelete", g.MarshalDelete)
	return v.Error()
}

var jsonEcd = binary.JSONEncoderDecoder{}

func marshalJSON[K gorp.Key, E gorp.Entry[K]](e E) ([]byte, error) {
	b, err := jsonEcd.Encode(context.TODO(), e)
	if err != nil {
		return nil, err
	}
	return append(b, '\n'), nil
}

// GorpConfigUUID is a helper function for creating a CDC pipeline that propagates
// changes to UUID keyed gorp entries written to the provided DB. The returned
// configuration should be passed to SubscribeToGorp.
func GorpConfigUUID[E gorp.Entry[uuid.UUID]](db *gorp.DB) GorpConfig[uuid.UUID, E] {
	return GorpConfig[uuid.UUID, E]{
		DB:             db,
		DeleteDataType: telem.UUIDT,
		SetDataType:    telem.JSONT,
		MarshalDelete:  func(k uuid.UUID) ([]byte, error) { return k[:], nil },
		MarshalSet:     marshalJSON[uuid.UUID, E],
	}
}

func GorpConfigString[E gorp.Entry[string]](db *gorp.DB) GorpConfig[string, E] {
	return GorpConfig[string, E]{
		DB:             db,
		DeleteDataType: telem.StringT,
		SetDataType:    telem.JSONT,
		MarshalDelete:  func(k string) ([]byte, error) { return append([]byte(k), '\n'), nil },
		MarshalSet:     marshalJSON[string, E],
	}
}

// SubscribeToGorp opens a CDC pipeline that subscribes to the sets and deletes of a
// particular entry type in the configured gorp.DB. The returned io.Closer should be
// closed to stop the CDC pipeline.
func SubscribeToGorp[K gorp.Key, E gorp.Entry[K]](
	ctx context.Context,
	svc *Provider,
	cfgs ...GorpConfig[K, E],
) (io.Closer, error) {
	cfg, err := config.New(DefaultGorpConfig[K, E](), cfgs...)
	if err != nil {
		return nil, err
	}
	var (
		obs = observe.Translator[gorp.TxReader[K, E], []change.Change[[]byte, struct{}]]{
			Observable: gorp.Observe[K, E](cfg.DB),
			Translate: func(r gorp.TxReader[K, E]) []change.Change[[]byte, struct{}] {
				out := make([]change.Change[[]byte, struct{}], 0, r.Count())
				for c, ok := r.Next(ctx); ok; c, ok = r.Next(ctx) {
					oc := change.Change[[]byte, struct{}]{Variant: c.Variant}
					if c.Variant == change.Set {
						v, err := cfg.MarshalSet(c.Value)
						if err != nil {
							svc.L.Error("failed to marshal set", zap.Error(err))
						}
						oc.Key = v
					} else {
						k, err := cfg.MarshalDelete(c.Key)
						if err != nil {
							svc.L.Error("failed to marshal delete", zap.Error(err))
						}
						oc.Key = k
					}
					out = append(out, oc)
				}
				return out
			},
		}
		obsCfg = ObservableConfig{
			Name:       fmt.Sprintf("gorp_%s", strings.ToLower(types.Name[E]())),
			Observable: obs,
			Set: channel.Channel{
				Name:     cfg.SetName,
				DataType: cfg.SetDataType,
			},
			Delete: channel.Channel{
				Name:     cfg.DeleteName,
				DataType: cfg.DeleteDataType,
			},
		}
	)
	return svc.SubscribeToObservable(ctx, obsCfg)
}
