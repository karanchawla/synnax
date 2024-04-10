// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package grpc

import (
	"github.com/synnaxlabs/freighter/fgrpc"
	"github.com/synnaxlabs/freighter/fnoop"
	"github.com/synnaxlabs/synnax/pkg/api"
	"go/types"
)

func New() (a api.Transport, transports []fgrpc.BindableTransport) {
	transports = make([]fgrpc.BindableTransport, 0, 20)
	transports = append(transports, newChannel(&a)...)
	transports = append(transports, newFramer(&a))
	transports = append(transports, newConnectivity(&a))
	transports = append(transports, newAuth(&a))
	transports = append(transports, newRanger(&a))
	transports = append(transports, newHardware(&a))

	// AUTH
	a.AuthChangeUsername = fnoop.UnaryServer[api.ChangeUsernameRequest, types.Nil]{}
	a.AuthChangePassword = fnoop.UnaryServer[api.ChangePasswordRequest, types.Nil]{}
	a.AuthRegistration = fnoop.UnaryServer[api.RegistrationRequest, api.TokenResponse]{}

	// RANGE
	a.RangeRename = fnoop.UnaryServer[api.RangeRenameRequest, types.Nil]{}

	// ONTOLOGY
	a.OntologyRetrieve = fnoop.UnaryServer[api.OntologyRetrieveRequest, api.OntologyRetrieveResponse]{}
	a.OntologyAddChildren = fnoop.UnaryServer[api.OntologyAddChildrenRequest, types.Nil]{}
	a.OntologyRemoveChildren = fnoop.UnaryServer[api.OntologyRemoveChildrenRequest, types.Nil]{}
	a.OntologyMoveChildren = fnoop.UnaryServer[api.OntologyMoveChildrenRequest, types.Nil]{}

	// GROUP
	a.OntologyGroupCreate = fnoop.UnaryServer[api.OntologyCreateGroupRequest, api.OntologyCreateGroupResponse]{}
	a.OntologyGroupDelete = fnoop.UnaryServer[api.OntologyDeleteGroupRequest, types.Nil]{}
	a.OntologyGroupRename = fnoop.UnaryServer[api.OntologyRenameGroupRequest, types.Nil]{}

	// WORKSPACE
	a.WorkspaceCreate = fnoop.UnaryServer[api.WorkspaceCreateRequest, api.WorkspaceCreateResponse]{}
	a.WorkspaceRetrieve = fnoop.UnaryServer[api.WorkspaceRetrieveRequest, api.WorkspaceRetrieveResponse]{}
	a.WorkspaceDelete = fnoop.UnaryServer[api.WorkspaceDeleteRequest, types.Nil]{}
	a.WorkspaceRename = fnoop.UnaryServer[api.WorkspaceRenameRequest, types.Nil]{}
	a.WorkspaceSetLayout = fnoop.UnaryServer[api.WorkspaceSetLayoutRequest, types.Nil]{}

	// PID
	a.PIDCreate = fnoop.UnaryServer[api.PIDCreateRequest, api.PIDCreateResponse]{}
	a.PIDDelete = fnoop.UnaryServer[api.PIDDeleteRequest, types.Nil]{}
	a.PIDRetrieve = fnoop.UnaryServer[api.PIDRetrieveRequest, api.PIDRetrieveResponse]{}
	a.PIDRename = fnoop.UnaryServer[api.PIDRenameRequest, types.Nil]{}
	a.PIDSetData = fnoop.UnaryServer[api.PIDSetDataRequest, types.Nil]{}
	a.PIDCopy = fnoop.UnaryServer[api.PIDCopyRequest, api.PIDCopyResponse]{}

	// LINE PLOT
	a.VisualizationCreate = fnoop.UnaryServer[api.VisualizationCreateRequest, api.VisualizationCreateResponse]{}
	a.VisualizationRetrieve = fnoop.UnaryServer[api.VisualizationRetrieveRequest, api.VisualizationRetrieveResponse]{}
	a.VisualizationDelete = fnoop.UnaryServer[api.VisualizationDeleteRequest, types.Nil]{}
	a.VisualizationRename = fnoop.UnaryServer[api.VisualizationRenameRequest, types.Nil]{}
	a.VisualizationSetData = fnoop.UnaryServer[api.VisualizationSetDataRequest, types.Nil]{}

	// LABEL
	a.LabelCreate = fnoop.UnaryServer[api.LabelCreateRequest, api.LabelCreateResponse]{}
	a.LabelRetrieve = fnoop.UnaryServer[api.LabelRetrieveRequest, api.LabelRetrieveResponse]{}
	a.LabelDelete = fnoop.UnaryServer[api.LabelDeleteRequest, types.Nil]{}
	a.LabelSet = fnoop.UnaryServer[api.LabelSetRequest, types.Nil]{}
	a.LabelRemove = fnoop.UnaryServer[api.LabelRemoveRequest, types.Nil]{}

	return a, transports
}
