// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/abc.proto

package abc

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

// AbcClient is the client API for Abc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AbcClient interface {
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Abc_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (Abc_ServerStreamClient, error)
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (Abc_BidiStreamClient, error)
}

type abcClient struct {
	cc grpc.ClientConnInterface
}

func NewAbcClient(cc grpc.ClientConnInterface) AbcClient {
	return &abcClient{cc}
}

func (c *abcClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/abc.Abc/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abcClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Abc_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Abc_ServiceDesc.Streams[0], "/abc.Abc/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &abcClientStreamClient{stream}
	return x, nil
}

type Abc_ClientStreamClient interface {
	Send(*ClientStreamRequest) error
	CloseAndRecv() (*ClientStreamResponse, error)
	grpc.ClientStream
}

type abcClientStreamClient struct {
	grpc.ClientStream
}

func (x *abcClientStreamClient) Send(m *ClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *abcClientStreamClient) CloseAndRecv() (*ClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *abcClient) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (Abc_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Abc_ServiceDesc.Streams[1], "/abc.Abc/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &abcServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Abc_ServerStreamClient interface {
	Recv() (*ServerStreamResponse, error)
	grpc.ClientStream
}

type abcServerStreamClient struct {
	grpc.ClientStream
}

func (x *abcServerStreamClient) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *abcClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (Abc_BidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Abc_ServiceDesc.Streams[2], "/abc.Abc/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &abcBidiStreamClient{stream}
	return x, nil
}

type Abc_BidiStreamClient interface {
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
	grpc.ClientStream
}

type abcBidiStreamClient struct {
	grpc.ClientStream
}

func (x *abcBidiStreamClient) Send(m *BidiStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *abcBidiStreamClient) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AbcServer is the server API for Abc service.
// All implementations must embed UnimplementedAbcServer
// for forward compatibility
type AbcServer interface {
	Call(context.Context, *CallRequest) (*CallResponse, error)
	ClientStream(Abc_ClientStreamServer) error
	ServerStream(*ServerStreamRequest, Abc_ServerStreamServer) error
	BidiStream(Abc_BidiStreamServer) error
	mustEmbedUnimplementedAbcServer()
}

// UnimplementedAbcServer must be embedded to have forward compatible implementations.
type UnimplementedAbcServer struct {
}

func (UnimplementedAbcServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedAbcServer) ClientStream(Abc_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedAbcServer) ServerStream(*ServerStreamRequest, Abc_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedAbcServer) BidiStream(Abc_BidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiStream not implemented")
}
func (UnimplementedAbcServer) mustEmbedUnimplementedAbcServer() {}

// UnsafeAbcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AbcServer will
// result in compilation errors.
type UnsafeAbcServer interface {
	mustEmbedUnimplementedAbcServer()
}

func RegisterAbcServer(s grpc.ServiceRegistrar, srv AbcServer) {
	s.RegisterService(&Abc_ServiceDesc, srv)
}

func _Abc_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbcServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abc.Abc/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbcServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Abc_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AbcServer).ClientStream(&abcClientStreamServer{stream})
}

type Abc_ClientStreamServer interface {
	SendAndClose(*ClientStreamResponse) error
	Recv() (*ClientStreamRequest, error)
	grpc.ServerStream
}

type abcClientStreamServer struct {
	grpc.ServerStream
}

func (x *abcClientStreamServer) SendAndClose(m *ClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *abcClientStreamServer) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Abc_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AbcServer).ServerStream(m, &abcServerStreamServer{stream})
}

type Abc_ServerStreamServer interface {
	Send(*ServerStreamResponse) error
	grpc.ServerStream
}

type abcServerStreamServer struct {
	grpc.ServerStream
}

func (x *abcServerStreamServer) Send(m *ServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Abc_BidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AbcServer).BidiStream(&abcBidiStreamServer{stream})
}

type Abc_BidiStreamServer interface {
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
	grpc.ServerStream
}

type abcBidiStreamServer struct {
	grpc.ServerStream
}

func (x *abcBidiStreamServer) Send(m *BidiStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *abcBidiStreamServer) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Abc_ServiceDesc is the grpc.ServiceDesc for Abc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Abc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abc.Abc",
	HandlerType: (*AbcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Abc_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _Abc_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _Abc_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidiStream",
			Handler:       _Abc_BidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/abc.proto",
}
