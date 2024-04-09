// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: synnax/pkg/api/grpc/v1/framer.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires grpc-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FrameIteratorService_Exec_FullMethodName = "/api.v1.FrameIteratorService/Exec"
)

// FrameIteratorServiceClient is the client API for FrameIteratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrameIteratorServiceClient interface {
	Exec(ctx context.Context, opts ...grpc.CallOption) (FrameIteratorService_ExecClient, error)
}

type frameIteratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFrameIteratorServiceClient(cc grpc.ClientConnInterface) FrameIteratorServiceClient {
	return &frameIteratorServiceClient{cc}
}

func (c *frameIteratorServiceClient) Exec(ctx context.Context, opts ...grpc.CallOption) (FrameIteratorService_ExecClient, error) {
	stream, err := c.cc.NewStream(ctx, &FrameIteratorService_ServiceDesc.Streams[0], FrameIteratorService_Exec_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &frameIteratorServiceExecClient{stream}
	return x, nil
}

type FrameIteratorService_ExecClient interface {
	Send(*FrameIteratorRequest) error
	Recv() (*FrameIteratorResponse, error)
	grpc.ClientStream
}

type frameIteratorServiceExecClient struct {
	grpc.ClientStream
}

func (x *frameIteratorServiceExecClient) Send(m *FrameIteratorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frameIteratorServiceExecClient) Recv() (*FrameIteratorResponse, error) {
	m := new(FrameIteratorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrameIteratorServiceServer is the server API for FrameIteratorService service.
// All implementations should embed UnimplementedFrameIteratorServiceServer
// for forward compatibility
type FrameIteratorServiceServer interface {
	Exec(FrameIteratorService_ExecServer) error
}

// UnimplementedFrameIteratorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFrameIteratorServiceServer struct {
}

func (UnimplementedFrameIteratorServiceServer) Exec(FrameIteratorService_ExecServer) error {
	return status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeFrameIteratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrameIteratorServiceServer will
// result in compilation errors.
type UnsafeFrameIteratorServiceServer interface {
	mustEmbedUnimplementedFrameIteratorServiceServer()
}

func RegisterFrameIteratorServiceServer(s grpc.ServiceRegistrar, srv FrameIteratorServiceServer) {
	s.RegisterService(&FrameIteratorService_ServiceDesc, srv)
}

func _FrameIteratorService_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrameIteratorServiceServer).Exec(&frameIteratorServiceExecServer{stream})
}

type FrameIteratorService_ExecServer interface {
	Send(*FrameIteratorResponse) error
	Recv() (*FrameIteratorRequest, error)
	grpc.ServerStream
}

type frameIteratorServiceExecServer struct {
	grpc.ServerStream
}

func (x *frameIteratorServiceExecServer) Send(m *FrameIteratorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frameIteratorServiceExecServer) Recv() (*FrameIteratorRequest, error) {
	m := new(FrameIteratorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrameIteratorService_ServiceDesc is the grpc.ServiceDesc for FrameIteratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FrameIteratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.FrameIteratorService",
	HandlerType: (*FrameIteratorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _FrameIteratorService_Exec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "synnax/pkg/api/grpc/v1/framer.proto",
}

const (
	FrameWriterService_Exec_FullMethodName = "/api.v1.FrameWriterService/Exec"
)

// FrameWriterServiceClient is the client API for FrameWriterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrameWriterServiceClient interface {
	Exec(ctx context.Context, opts ...grpc.CallOption) (FrameWriterService_ExecClient, error)
}

type frameWriterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFrameWriterServiceClient(cc grpc.ClientConnInterface) FrameWriterServiceClient {
	return &frameWriterServiceClient{cc}
}

func (c *frameWriterServiceClient) Exec(ctx context.Context, opts ...grpc.CallOption) (FrameWriterService_ExecClient, error) {
	stream, err := c.cc.NewStream(ctx, &FrameWriterService_ServiceDesc.Streams[0], FrameWriterService_Exec_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &frameWriterServiceExecClient{stream}
	return x, nil
}

type FrameWriterService_ExecClient interface {
	Send(*FrameWriterRequest) error
	Recv() (*FrameWriterResponse, error)
	grpc.ClientStream
}

type frameWriterServiceExecClient struct {
	grpc.ClientStream
}

func (x *frameWriterServiceExecClient) Send(m *FrameWriterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frameWriterServiceExecClient) Recv() (*FrameWriterResponse, error) {
	m := new(FrameWriterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrameWriterServiceServer is the server API for FrameWriterService service.
// All implementations should embed UnimplementedFrameWriterServiceServer
// for forward compatibility
type FrameWriterServiceServer interface {
	Exec(FrameWriterService_ExecServer) error
}

// UnimplementedFrameWriterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFrameWriterServiceServer struct {
}

func (UnimplementedFrameWriterServiceServer) Exec(FrameWriterService_ExecServer) error {
	return status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeFrameWriterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrameWriterServiceServer will
// result in compilation errors.
type UnsafeFrameWriterServiceServer interface {
	mustEmbedUnimplementedFrameWriterServiceServer()
}

func RegisterFrameWriterServiceServer(s grpc.ServiceRegistrar, srv FrameWriterServiceServer) {
	s.RegisterService(&FrameWriterService_ServiceDesc, srv)
}

func _FrameWriterService_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrameWriterServiceServer).Exec(&frameWriterServiceExecServer{stream})
}

type FrameWriterService_ExecServer interface {
	Send(*FrameWriterResponse) error
	Recv() (*FrameWriterRequest, error)
	grpc.ServerStream
}

type frameWriterServiceExecServer struct {
	grpc.ServerStream
}

func (x *frameWriterServiceExecServer) Send(m *FrameWriterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frameWriterServiceExecServer) Recv() (*FrameWriterRequest, error) {
	m := new(FrameWriterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrameWriterService_ServiceDesc is the grpc.ServiceDesc for FrameWriterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FrameWriterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.FrameWriterService",
	HandlerType: (*FrameWriterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _FrameWriterService_Exec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "synnax/pkg/api/grpc/v1/framer.proto",
}

const (
	FrameStreamerService_Exec_FullMethodName = "/api.v1.FrameStreamerService/Exec"
)

// FrameStreamerServiceClient is the client API for FrameStreamerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrameStreamerServiceClient interface {
	Exec(ctx context.Context, opts ...grpc.CallOption) (FrameStreamerService_ExecClient, error)
}

type frameStreamerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFrameStreamerServiceClient(cc grpc.ClientConnInterface) FrameStreamerServiceClient {
	return &frameStreamerServiceClient{cc}
}

func (c *frameStreamerServiceClient) Exec(ctx context.Context, opts ...grpc.CallOption) (FrameStreamerService_ExecClient, error) {
	stream, err := c.cc.NewStream(ctx, &FrameStreamerService_ServiceDesc.Streams[0], FrameStreamerService_Exec_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &frameStreamerServiceExecClient{stream}
	return x, nil
}

type FrameStreamerService_ExecClient interface {
	Send(*FrameStreamerRequest) error
	Recv() (*FrameStreamerResponse, error)
	grpc.ClientStream
}

type frameStreamerServiceExecClient struct {
	grpc.ClientStream
}

func (x *frameStreamerServiceExecClient) Send(m *FrameStreamerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frameStreamerServiceExecClient) Recv() (*FrameStreamerResponse, error) {
	m := new(FrameStreamerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrameStreamerServiceServer is the server API for FrameStreamerService service.
// All implementations should embed UnimplementedFrameStreamerServiceServer
// for forward compatibility
type FrameStreamerServiceServer interface {
	Exec(FrameStreamerService_ExecServer) error
}

// UnimplementedFrameStreamerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFrameStreamerServiceServer struct {
}

func (UnimplementedFrameStreamerServiceServer) Exec(FrameStreamerService_ExecServer) error {
	return status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

// UnsafeFrameStreamerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrameStreamerServiceServer will
// result in compilation errors.
type UnsafeFrameStreamerServiceServer interface {
	mustEmbedUnimplementedFrameStreamerServiceServer()
}

func RegisterFrameStreamerServiceServer(s grpc.ServiceRegistrar, srv FrameStreamerServiceServer) {
	s.RegisterService(&FrameStreamerService_ServiceDesc, srv)
}

func _FrameStreamerService_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrameStreamerServiceServer).Exec(&frameStreamerServiceExecServer{stream})
}

type FrameStreamerService_ExecServer interface {
	Send(*FrameStreamerResponse) error
	Recv() (*FrameStreamerRequest, error)
	grpc.ServerStream
}

type frameStreamerServiceExecServer struct {
	grpc.ServerStream
}

func (x *frameStreamerServiceExecServer) Send(m *FrameStreamerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frameStreamerServiceExecServer) Recv() (*FrameStreamerRequest, error) {
	m := new(FrameStreamerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrameStreamerService_ServiceDesc is the grpc.ServiceDesc for FrameStreamerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FrameStreamerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.FrameStreamerService",
	HandlerType: (*FrameStreamerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _FrameStreamerService_Exec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "synnax/pkg/api/grpc/v1/framer.proto",
}