package handler

import (
	"context"
	"io"
	"time"

	pb "abc/proto"
)

type Abc struct {
	pb.UnimplementedAbcServer
}

func (e *Abc) Call(ctx context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	return &pb.CallResponse{
		Msg: "Hello " + req.Name,
	}, nil
}

func (e *Abc) ClientStream(stream pb.Abc_ClientStreamServer) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		count++
	}
}

func (e *Abc) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Abc_ServerStreamServer) error {
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Abc) BidiStream(stream pb.Abc_BidiStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
