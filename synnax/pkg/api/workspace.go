// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package api

import (
	"context"
	"github.com/google/uuid"
	"github.com/synnaxlabs/synnax/pkg/api/errors"
	"github.com/synnaxlabs/synnax/pkg/user"
	"github.com/synnaxlabs/synnax/pkg/workspace"
	"github.com/synnaxlabs/synnax/pkg/workspace/vis"
	"github.com/synnaxlabs/x/gorp"
	"go/types"
)

type WorkspaceService struct {
	dbProvider
	internal *workspace.Service
}

func NewWorkspaceService(p Provider) *WorkspaceService {
	return &WorkspaceService{
		dbProvider: p.db,
		internal:   p.Config.Workspace,
	}
}

type WorkspaceCreateRequest struct {
	Workspaces []workspace.Workspace `json:"workspaces" msgpack:"workspaces"`
}

type WorkspaceCreateResponse struct {
	Workspaces []workspace.Workspace `json:"workspaces" msgpack:"workspaces"`
}

func (s *WorkspaceService) Create(ctx context.Context, req WorkspaceCreateRequest) (res WorkspaceCreateResponse, err error) {
	userKey, err_ := user.FromOntologyID(getSubject(ctx))
	if err_ != nil {
		return res, errors.Unexpected(err_)
	}
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		for _, w := range req.Workspaces {
			w.Author = userKey
			err := s.internal.NewWriter(tx).Create(ctx, &w)
			if err != nil {
				return err
			}
			res.Workspaces = append(res.Workspaces, w)
		}
		return nil
	})
}

type WorkspaceRenameRequest struct {
	Key  uuid.UUID `json:"key" msgpack:"key"`
	Name string    `json:"name" msgpack:"name"`
}

func (s *WorkspaceService) Rename(ctx context.Context, req WorkspaceRenameRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).Rename(ctx, req.Key, req.Name)
	})
}

type WorkspaceSetLayoutRequest struct {
	Key    uuid.UUID `json:"key" msgpack:"key"`
	Layout string    `json:"layout" msgpack:"layout"`
}

func (s *WorkspaceService) SetLayout(ctx context.Context, req WorkspaceSetLayoutRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).SetLayout(ctx, req.Key, req.Layout)
	})
}

type WorkspaceRetrieveRequest struct {
	Keys   []uuid.UUID `json:"keys" msgpack:"keys"`
	Search string      `json:"search" msgpack:"search"`
	Author uuid.UUID   `json:"author" msgpack:"author"`
	Limit  int         `json:"limit" msgpack:"limit"`
	Offset int         `json:"offset" msgpack:"offset"`
}

type WorkspaceRetrieveResponse struct {
	Workspaces []workspace.Workspace `json:"workspaces" msgpack:"workspaces"`
}

func (s *WorkspaceService) Retrieve(
	ctx context.Context,
	req WorkspaceRetrieveRequest,
) (res WorkspaceRetrieveResponse, err error) {
	q := s.internal.NewRetrieve().Search(req.Search)
	if len(req.Keys) > 0 {
		q = q.WhereKeys(req.Keys...)
	}
	if req.Author != uuid.Nil {
		q = q.WhereAuthor(req.Author)
	}
	if req.Limit > 0 {
		q = q.Limit(req.Limit)
	}
	if req.Offset > 0 {
		q = q.Offset(req.Offset)
	}
	err = q.Entries(&res.Workspaces).Exec(ctx, nil)
	return res, err
}

type WorkspaceDeleteRequest struct {
	Keys []uuid.UUID `json:"keys" msgpack:"keys"`
}

func (s *WorkspaceService) Delete(ctx context.Context, req WorkspaceDeleteRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).Delete(ctx, req.Keys...)
	})
}

type VisService struct {
	dbProvider
	internal *vis.Service
}

func NewVisService(p Provider) *VisService {
	return &VisService{
		dbProvider: p.db,
		internal:   p.Config.Vis,
	}
}

type VisCreateRequest struct {
	Workspace uuid.UUID `json:"workspace" msgpack:"workspace"`
	Viss      []vis.Vis `json:"line_plots" msgpack:"line_plots"`
}

type VisCreateResponse struct {
	Viss []vis.Vis `json:"viss" msgpack:"viss"`
}

func (s *VisService) Create(ctx context.Context, req VisCreateRequest) (res VisCreateResponse, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		for _, lp := range req.Viss {
			err := s.internal.NewWriter(tx).Create(ctx, req.Workspace, &lp)
			if err != nil {
				return err
			}
			res.Viss = append(res.Viss, lp)
		}
		return err
	})
}

type VisRenameRequest struct {
	Key  uuid.UUID `json:"key" msgpack:"key"`
	Name string    `json:"name" msgpack:"name"`
}

func (s *VisService) Rename(ctx context.Context, req VisRenameRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).Rename(ctx, req.Key, req.Name)
	})
}

type VisSetDataRequest struct {
	Key  uuid.UUID `json:"key" msgpack:"key"`
	Data string    `json:"data" msgpack:"data"`
}

func (s *VisService) SetData(ctx context.Context, req VisSetDataRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).SetData(ctx, req.Key, req.Data)
	})
}

type VisRetrieveRequest struct {
	Keys []uuid.UUID `json:"keys" msgpack:"keys"`
}

type VisRetrieveResponse struct {
	Vis []vis.Vis `json:"vis" msgpack:"vis"`
}

func (s *VisService) Retrieve(ctx context.Context, req VisRetrieveRequest) (res VisRetrieveResponse, err error) {
	err = s.internal.NewRetrieve().
		WhereKeys(req.Keys...).Entries(&res.Vis).Exec(ctx, nil)
	return res, err
}

type VisDeleteRequest struct {
	Keys []uuid.UUID `json:"keys" msgpack:"keys"`
}

func (s *VisService) Delete(ctx context.Context, req VisDeleteRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).Delete(ctx, req.Keys...)
	})
}

type VisCopyRequest struct {
	Key      uuid.UUID `json:"key" msgpack:"key"`
	Name     string    `json:"name" msgpack:"name"`
	Snapshot bool      `json:"snapshot" msgpack:"snapshot"`
}

type VisCopyResponse struct {
	Vis vis.Vis `json:"vis" msgpack:"vis"`
}

func (s *VisService) Copy(ctx context.Context, req VisCopyRequest) (res VisCopyResponse, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return errors.Auto(s.internal.NewWriter(tx).Copy(ctx, req.Key, req.Name, req.Snapshot, &res.Vis))
	})
}
