// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package http

import (
	"go/types"

	"github.com/synnaxlabs/freighter/fhttp"
	"github.com/synnaxlabs/synnax/pkg/api"
	"github.com/synnaxlabs/synnax/pkg/auth"
)

func New(router *fhttp.Router) (t api.Transport) {
	// AUTH
	t.AuthLogin = fhttp.UnaryServer[auth.InsecureCredentials, api.TokenResponse](router, "/api/v1/auth/login")
	t.AuthRegistration = fhttp.UnaryServer[api.RegistrationRequest, api.TokenResponse](router, "/api/v1/auth/register")
	t.AuthChangePassword = fhttp.UnaryServer[api.ChangePasswordRequest, types.Nil](router, "/api/v1/auth/protected/change-password")
	t.AuthChangeUsername = fhttp.UnaryServer[api.ChangeUsernameRequest, types.Nil](router, "/api/v1/auth/protected/change-username")

	// CHANNEL
	t.ChannelCreate = fhttp.UnaryServer[api.ChannelCreateRequest, api.ChannelCreateResponse](router, "/api/v1/channel/create")
	t.ChannelRetrieve = fhttp.UnaryServer[api.ChannelRetrieveRequest, api.ChannelRetrieveResponse](router, "/api/v1/channel/retrieve")
	t.ConnectivityCheck = fhttp.UnaryServer[types.Nil, api.ConnectivityCheckResponse](router, "/api/v1/connectivity/check")

	// FRAME
	t.FrameWriter = fhttp.StreamServer[api.FrameWriterRequest, api.FrameWriterResponse](router, "/api/v1/frame/write")
	t.FrameIterator = fhttp.StreamServer[api.FrameIteratorRequest, api.FrameIteratorResponse](router, "/api/v1/frame/iterate")
	t.FrameStreamer = fhttp.StreamServer[api.FrameStreamerRequest, api.FrameStreamerResponse](router, "/api/v1/frame/stream")

	// ONTOLOGY
	t.OntologyRetrieve = fhttp.UnaryServer[api.OntologyRetrieveRequest, api.OntologyRetrieveResponse](router, "/api/v1/ontology/retrieve")
	t.OntologyAddChildren = fhttp.UnaryServer[api.OntologyAddChildrenRequest, types.Nil](router, "/api/v1/ontology/add-children")
	t.OntologyRemoveChildren = fhttp.UnaryServer[api.OntologyRemoveChildrenRequest, types.Nil](router, "/api/v1/ontology/remove-children")
	t.OntologyMoveChildren = fhttp.UnaryServer[api.OntologyMoveChildrenRequest, types.Nil](router, "/api/v1/ontology/move-children")

	// GROUP
	t.OntologyGroupCreate = fhttp.UnaryServer[api.OntologyCreateGroupRequest, api.OntologyCreateGroupResponse](router, "/api/v1/ontology/create-group")
	t.OntologyGroupDelete = fhttp.UnaryServer[api.OntologyDeleteGroupRequest, types.Nil](router, "/api/v1/ontology/delete-group")
	t.OntologyGroupRename = fhttp.UnaryServer[api.OntologyRenameGroupRequest, types.Nil](router, "/api/v1/ontology/rename-group")

	// RANGE
	t.RangeRetrieve = fhttp.UnaryServer[api.RangeRetrieveRequest, api.RangeRetrieveResponse](router, "/api/v1/range/retrieve")
	t.RangeCreate = fhttp.UnaryServer[api.RangeCreateRequest, api.RangeCreateResponse](router, "/api/v1/range/create")
	t.RangeDelete = fhttp.UnaryServer[api.RangeDeleteRequest, types.Nil](router, "/api/v1/range/delete")
	t.RangeRename = fhttp.UnaryServer[api.RangeRenameRequest, types.Nil](router, "/api/v1/range/rename")
	t.RangeKVGet = fhttp.UnaryServer[api.RangeKVGetRequest, api.RangeKVGetResponse](router, "/api/v1/range/kv/get")
	t.RangeKVSet = fhttp.UnaryServer[api.RangeKVSetRequest, types.Nil](router, "/api/v1/range/kv/set")
	t.RangeKVDelete = fhttp.UnaryServer[api.RangeKVDeleteRequest, types.Nil](router, "/api/v1/range/kv/delete")
	t.RangeAliasSet = fhttp.UnaryServer[api.RangeAliasSetRequest, types.Nil](router, "/api/v1/range/alias/set")
	t.RangeAliasResolve = fhttp.UnaryServer[api.RangeAliasResolveRequest, api.RangeAliasResolveResponse](router, "/api/v1/range/alias/resolve")
	t.RangeAliasList = fhttp.UnaryServer[api.RangeAliasListRequest, api.RangeAliasListResponse](router, "/api/v1/range/alias/list")
	t.RangeAliasDelete = fhttp.UnaryServer[api.RangeAliasDeleteRequest, types.Nil](router, "/api/v1/range/alias/delete")
	t.RangeSetActive = fhttp.UnaryServer[api.RangeSetActiveRequest, types.Nil](router, "/api/v1/range/set-active")
	t.RangeRetrieveActive = fhttp.UnaryServer[types.Nil, api.RangeRetrieveActiveResponse](router, "/api/v1/range/retrieve-active")
	t.RangeClearActive = fhttp.UnaryServer[types.Nil, types.Nil](router, "/api/v1/range/clear-active")

	// WORKSPACE
	t.WorkspaceCreate = fhttp.UnaryServer[api.WorkspaceCreateRequest, api.WorkspaceCreateResponse](router, "/api/v1/workspace/create")
	t.WorkspaceRetrieve = fhttp.UnaryServer[api.WorkspaceRetrieveRequest, api.WorkspaceRetrieveResponse](router, "/api/v1/workspace/retrieve")
	t.WorkspaceDelete = fhttp.UnaryServer[api.WorkspaceDeleteRequest, types.Nil](router, "/api/v1/workspace/delete")
	t.WorkspaceRename = fhttp.UnaryServer[api.WorkspaceRenameRequest, types.Nil](router, "/api/v1/workspace/rename")
	t.WorkspaceSetLayout = fhttp.UnaryServer[api.WorkspaceSetLayoutRequest, types.Nil](router, "/api/v1/workspace/set-layout")

	// Vis
	t.VisCreate = fhttp.UnaryServer[api.VisCreateRequest, api.VisCreateResponse](router, "/api/v1/workspace/vis/create")
	t.VisRetrieve = fhttp.UnaryServer[api.VisRetrieveRequest, api.VisRetrieveResponse](router, "/api/v1/workspace/vis/retrieve")
	t.VisDelete = fhttp.UnaryServer[api.VisDeleteRequest, types.Nil](router, "/api/v1/workspace/vis/delete")
	t.VisRename = fhttp.UnaryServer[api.VisRenameRequest, types.Nil](router, "/api/v1/workspace/vis/rename")
	t.VisSetData = fhttp.UnaryServer[api.VisSetDataRequest, types.Nil](router, "/api/v1/workspace/vis/set-data")
	t.VisCopy = fhttp.UnaryServer[api.VisCopyRequest, api.VisCopyResponse](router, "/api/v1/workspace/vis/copy")

	// LABEL
	t.LabelCreate = fhttp.UnaryServer[api.LabelCreateRequest, api.LabelCreateResponse](router, "/api/v1/label/create")
	t.LabelRetrieve = fhttp.UnaryServer[api.LabelRetrieveRequest, api.LabelRetrieveResponse](router, "/api/v1/label/retrieve")
	t.LabelDelete = fhttp.UnaryServer[api.LabelDeleteRequest, types.Nil](router, "/api/v1/label/delete")
	t.LabelSet = fhttp.UnaryServer[api.LabelSetRequest, types.Nil](router, "/api/v1/label/set")
	t.LabelRemove = fhttp.UnaryServer[api.LabelRemoveRequest, types.Nil](router, "/api/v1/label/remove")

	// DEVICE
	t.HardwareCreateRack = fhttp.UnaryServer[api.HardwareCreateRackRequest, api.HardwareCreateRackResponse](router, "/api/v1/hardware/rack/create")
	t.HardwareRetrieveRack = fhttp.UnaryServer[api.HardwareRetrieveRackRequest, api.HardwareRetrieveRackResponse](router, "/api/v1/hardware/rack/retrieve")
	t.HardwareDeleteRack = fhttp.UnaryServer[api.HardwareDeleteRackRequest, types.Nil](router, "/api/v1/hardware/rack/delete")
	t.HardwareCreateTask = fhttp.UnaryServer[api.HardwareCreateTaskRequest, api.HardwareCreateTaskResponse](router, "/api/v1/hardware/task/create")
	t.HardwareRetrieveTask = fhttp.UnaryServer[api.HardwareRetrieveTaskRequest, api.HardwareRetrieveTaskResponse](router, "/api/v1/hardware/task/retrieve")
	t.HardwareDeleteTask = fhttp.UnaryServer[api.HardwareDeleteTaskRequest, types.Nil](router, "/api/v1/hardware/task/delete")
	t.HardwareCreateDevice = fhttp.UnaryServer[api.HardwareCreateDeviceRequest, api.HardwareCreateDeviceResponse](router, "/api/v1/hardware/device/create")
	t.HardwareRetrieveDevice = fhttp.UnaryServer[api.HardwareRetrieveDeviceRequest, api.HardwareRetrieveDeviceResponse](router, "/api/v1/hardware/device/retrieve")
	t.HardwareDeleteDevice = fhttp.UnaryServer[api.HardwareDeleteDeviceRequest, types.Nil](router, "/api/v1/hardware/device/delete")

	return t
}
