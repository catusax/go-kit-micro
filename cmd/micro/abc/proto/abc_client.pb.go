// Code generated by protoc-gen-go-ascii. DO NOT EDIT.

package abc

import (
	"context"
	"github.com/coolrc136/go-kit-micro/sd"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

type ClientImpl struct {
	call         endpoint.Endpoint
	clientStream endpoint.Endpoint
	serverStream endpoint.Endpoint
	bidiStream   endpoint.Endpoint
}

func NewAbcClientImpl(name string, logger log.Logger) *ClientImpl {
	instancer, err := sd.NewInstancer(name, logger)
	if err != nil {
		panic(err)
	}

	return &ClientImpl{

		call: sd.GetEndPoint(instancer, func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error) {
			client := NewAbcClient(conn)
			req := request.(*CallRequest)

			return client.Call(ctx, req)
		}, logger),
		clientStream: sd.GetEndPoint(instancer, func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error) {
			client := NewAbcClient(conn)
			req := request.(*ClientStreamRequest)

			return client.ClientStream(ctx, req)
		}, logger),
		serverStream: sd.GetEndPoint(instancer, func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error) {
			client := NewAbcClient(conn)
			req := request.(*ServerStreamRequest)

			return client.ServerStream(ctx, req)
		}, logger),
		bidiStream: sd.GetEndPoint(instancer, func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error) {
			client := NewAbcClient(conn)
			req := request.(*BidiStreamRequest)

			return client.BidiStream(ctx, req)
		}, logger),
	}
}

// Call
func (n *ClientImpl) Call(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	rsp, err := n.call(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rsp.(*CallResponse)
	return res, err
}

// ClientStream
func (n *ClientImpl) ClientStream(ctx context.Context, req *ClientStreamRequest) (*ClientStreamResponse, error) {
	rsp, err := n.clientStream(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rsp.(*ClientStreamResponse)
	return res, err
}

// ServerStream
func (n *ClientImpl) ServerStream(ctx context.Context, req *ServerStreamRequest) (*ServerStreamResponse, error) {
	rsp, err := n.serverStream(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rsp.(*ServerStreamResponse)
	return res, err
}

// BidiStream
func (n *ClientImpl) BidiStream(ctx context.Context, req *BidiStreamRequest) (*BidiStreamResponse, error) {
	rsp, err := n.bidiStream(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rsp.(*BidiStreamResponse)
	return res, err
}
