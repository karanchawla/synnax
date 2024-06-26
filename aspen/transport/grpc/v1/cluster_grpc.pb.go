// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aspen/transport/grpc/v1/cluster.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClusterGossipService_Exec_FullMethodName = "/aspen.v1.ClusterGossipService/Exec"
)

// ClusterGossipServiceClient is the client API for ClusterGossipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterGossipServiceClient interface {
	Exec(ctx context.Context, in *ClusterGossip, opts ...grpc.CallOption) (*ClusterGossip, error)
}

type clusterGossipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterGossipServiceClient(cc grpc.ClientConnInterface) ClusterGossipServiceClient {
	return &clusterGossipServiceClient{cc}
}

func (c *clusterGossipServiceClient) Exec(ctx context.Context, in *ClusterGossip, opts ...grpc.CallOption) (*ClusterGossip, error) {
	out := new(ClusterGossip)
	err := c.cc.Invoke(ctx, ClusterGossipService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterGossipServiceServer is the server API for ClusterGossipService service.
// All implementations should embed UnimplementedClusterGossipServiceServer
// for forward compatibility
type ClusterGossipServiceServer interface {
	Exec(context.Context, *ClusterGossip) (*ClusterGossip, error)
}

// UnimplementedClusterGossipServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterGossipServiceServer struct {
}

func (UnimplementedClusterGossipServiceServer) Exec(context.Context, *ClusterGossip) (*ClusterGossip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeClusterGossipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterGossipServiceServer will
// result in compilation errors.
type UnsafeClusterGossipServiceServer interface {
	mustEmbedUnimplementedClusterGossipServiceServer()
}

func RegisterClusterGossipServiceServer(s grpc.ServiceRegistrar, srv ClusterGossipServiceServer) {
	s.RegisterService(&ClusterGossipService_ServiceDesc, srv)
}

func _ClusterGossipService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterGossip)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterGossipServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterGossipService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterGossipServiceServer).Exec(ctx, req.(*ClusterGossip))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterGossipService_ServiceDesc is the grpc.ServiceDesc for ClusterGossipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterGossipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aspen.v1.ClusterGossipService",
	HandlerType: (*ClusterGossipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _ClusterGossipService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aspen/transport/grpc/v1/cluster.proto",
}

const (
	PledgeService_Exec_FullMethodName = "/aspen.v1.PledgeService/Exec"
)

// PledgeServiceClient is the client API for PledgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PledgeServiceClient interface {
	Exec(ctx context.Context, in *ClusterPledge, opts ...grpc.CallOption) (*ClusterPledge, error)
}

type pledgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPledgeServiceClient(cc grpc.ClientConnInterface) PledgeServiceClient {
	return &pledgeServiceClient{cc}
}

func (c *pledgeServiceClient) Exec(ctx context.Context, in *ClusterPledge, opts ...grpc.CallOption) (*ClusterPledge, error) {
	out := new(ClusterPledge)
	err := c.cc.Invoke(ctx, PledgeService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PledgeServiceServer is the server API for PledgeService service.
// All implementations should embed UnimplementedPledgeServiceServer
// for forward compatibility
type PledgeServiceServer interface {
	Exec(context.Context, *ClusterPledge) (*ClusterPledge, error)
}

// UnimplementedPledgeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPledgeServiceServer struct {
}

func (UnimplementedPledgeServiceServer) Exec(context.Context, *ClusterPledge) (*ClusterPledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafePledgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PledgeServiceServer will
// result in compilation errors.
type UnsafePledgeServiceServer interface {
	mustEmbedUnimplementedPledgeServiceServer()
}

func RegisterPledgeServiceServer(s grpc.ServiceRegistrar, srv PledgeServiceServer) {
	s.RegisterService(&PledgeService_ServiceDesc, srv)
}

func _PledgeService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterPledge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PledgeServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PledgeService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PledgeServiceServer).Exec(ctx, req.(*ClusterPledge))
	}
	return interceptor(ctx, in, info, handler)
}

// PledgeService_ServiceDesc is the grpc.ServiceDesc for PledgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PledgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aspen.v1.PledgeService",
	HandlerType: (*PledgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _PledgeService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aspen/transport/grpc/v1/cluster.proto",
}
