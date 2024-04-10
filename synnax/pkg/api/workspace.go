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

type VisualizationService struct {
	dbProvider
	internal *vis.Service
}

func NewVisualizationService(p Provider) *VisualizationService {
	return &VisualizationService{
		dbProvider: p.db,
		internal:   p.Config.Visualization,
	}
}

type VisualizationCreateRequest struct {
	Workspace      uuid.UUID `json:"workspace" msgpack:"workspace"`
	Visualizations []vis.Vis `json:"line_plots" msgpack:"line_plots"`
}

type VisualizationCreateResponse struct {
	Visualizations []vis.Vis `json:"visualizations" msgpack:"visualizations"`
}

func (s *VisualizationService) Create(ctx context.Context, req VisualizationCreateRequest) (res VisualizationCreateResponse, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		for _, lp := range req.Visualizations {
			err := s.internal.NewWriter(tx).Create(ctx, req.Workspace, &lp)
			if err != nil {
				return err
			}
			res.Visualizations = append(res.Visualizations, lp)
		}
		return err
	})
}

type VisualizationRenameRequest struct {
	Key  uuid.UUID `json:"key" msgpack:"key"`
	Name string    `json:"name" msgpack:"name"`
}

func (s *VisualizationService) Rename(ctx context.Context, req VisualizationRenameRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).Rename(ctx, req.Key, req.Name)
	})
}

type VisualizationSetDataRequest struct {
	Key  uuid.UUID `json:"key" msgpack:"key"`
	Data string    `json:"data" msgpack:"data"`
}

func (s *VisualizationService) SetData(ctx context.Context, req VisualizationSetDataRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).SetData(ctx, req.Key, req.Data)
	})
}

type VisualizationRetrieveRequest struct {
	Keys []uuid.UUID `json:"keys" msgpack:"keys"`
}

type VisualizationRetrieveResponse struct {
	Visualizations []vis.Vis `json:"line_plots" msgpack:"line_plots"`
}

func (s *VisualizationService) Retrieve(ctx context.Context, req VisualizationRetrieveRequest) (res VisualizationRetrieveResponse, err error) {
	err = s.internal.NewRetrieve().
		WhereKeys(req.Keys...).Entries(&res.Visualizations).Exec(ctx, nil)
	return res, err
}

type VisualizationDeleteRequest struct {
	Keys []uuid.UUID `json:"keys" msgpack:"keys"`
}

func (s *VisualizationService) Delete(ctx context.Context, req VisualizationDeleteRequest) (res types.Nil, err error) {
	return res, s.WithTx(ctx, func(tx gorp.Tx) error {
		return s.internal.NewWriter(tx).Delete(ctx, req.Keys...)
	})
}
