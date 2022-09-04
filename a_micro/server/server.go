package server

import (
	"context"
	pb "grpc_micro/a_micro/a_protopb"
	"grpc_micro/b_micro"
)

type AServer struct {
	Bc b_micro.Client
}

func (s *AServer) GetA(ctx context.Context, in *pb.ARequest) (*pb.AResponse, error) {
	res, err := s.Bc.GetB(ctx, in.ParamB)
	if err != nil {
		return nil, err
	}

	return &pb.AResponse{
		ResA: res.ResA,
		ResB: res.ResB,
	}, nil
}

func NewAServer(bCleint b_micro.Client) *AServer {
	return &AServer{Bc: bCleint}
}
