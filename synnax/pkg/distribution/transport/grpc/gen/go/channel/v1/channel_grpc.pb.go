// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: channel/v1/channel.proto

package channelv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChannelCreateService_Exec_FullMethodName = "/channel.v1.ChannelCreateService/Exec"
)

// ChannelCreateServiceClient is the client API for ChannelCreateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelCreateServiceClient interface {
	Exec(ctx context.Context, in *CreateMessage, opts ...grpc.CallOption) (*CreateMessage, error)
}

type channelCreateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelCreateServiceClient(cc grpc.ClientConnInterface) ChannelCreateServiceClient {
	return &channelCreateServiceClient{cc}
}

func (c *channelCreateServiceClient) Exec(ctx context.Context, in *CreateMessage, opts ...grpc.CallOption) (*CreateMessage, error) {
	out := new(CreateMessage)
	err := c.cc.Invoke(ctx, ChannelCreateService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelCreateServiceServer is the server API for ChannelCreateService service.
// All implementations should embed UnimplementedChannelCreateServiceServer
// for forward compatibility
type ChannelCreateServiceServer interface {
	Exec(context.Context, *CreateMessage) (*CreateMessage, error)
}

// UnimplementedChannelCreateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChannelCreateServiceServer struct {
}

func (UnimplementedChannelCreateServiceServer) Exec(context.Context, *CreateMessage) (*CreateMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeChannelCreateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelCreateServiceServer will
// result in compilation errors.
type UnsafeChannelCreateServiceServer interface {
	mustEmbedUnimplementedChannelCreateServiceServer()
}

func RegisterChannelCreateServiceServer(s grpc.ServiceRegistrar, srv ChannelCreateServiceServer) {
	s.RegisterService(&ChannelCreateService_ServiceDesc, srv)
}

func _ChannelCreateService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelCreateServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelCreateService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelCreateServiceServer).Exec(ctx, req.(*CreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelCreateService_ServiceDesc is the grpc.ServiceDesc for ChannelCreateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelCreateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.v1.ChannelCreateService",
	HandlerType: (*ChannelCreateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _ChannelCreateService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/v1/channel.proto",
}

const (
	ChannelDeleteService_Exec_FullMethodName = "/channel.v1.ChannelDeleteService/Exec"
)

// ChannelDeleteServiceClient is the client API for ChannelDeleteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelDeleteServiceClient interface {
	Exec(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type channelDeleteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelDeleteServiceClient(cc grpc.ClientConnInterface) ChannelDeleteServiceClient {
	return &channelDeleteServiceClient{cc}
}

func (c *channelDeleteServiceClient) Exec(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChannelDeleteService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelDeleteServiceServer is the server API for ChannelDeleteService service.
// All implementations should embed UnimplementedChannelDeleteServiceServer
// for forward compatibility
type ChannelDeleteServiceServer interface {
	Exec(context.Context, *DeleteRequest) (*emptypb.Empty, error)
}

// UnimplementedChannelDeleteServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChannelDeleteServiceServer struct {
}

func (UnimplementedChannelDeleteServiceServer) Exec(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeChannelDeleteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelDeleteServiceServer will
// result in compilation errors.
type UnsafeChannelDeleteServiceServer interface {
	mustEmbedUnimplementedChannelDeleteServiceServer()
}

func RegisterChannelDeleteServiceServer(s grpc.ServiceRegistrar, srv ChannelDeleteServiceServer) {
	s.RegisterService(&ChannelDeleteService_ServiceDesc, srv)
}

func _ChannelDeleteService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelDeleteServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelDeleteService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelDeleteServiceServer).Exec(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelDeleteService_ServiceDesc is the grpc.ServiceDesc for ChannelDeleteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelDeleteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.v1.ChannelDeleteService",
	HandlerType: (*ChannelDeleteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _ChannelDeleteService_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/v1/channel.proto",
}