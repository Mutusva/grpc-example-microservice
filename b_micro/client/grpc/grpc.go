package grpc

import (
	"context"
	"google.golang.org/grpc"
	"grpc_micro/b_micro"
	pb "grpc_micro/b_micro/b_protopb"
)

var _ b_micro.Client = (*bClient)(nil)

type bClient struct {
	rpcConn     *grpc.ClientConn
	bClientConn pb.BMicroserviceClient
}

func (bc *bClient) GetB(ctx context.Context, param int64) (*b_micro.Bres, error) {
	in := pb.BRequest{ParamA: param}
	res, err := bc.bClientConn.GetB(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &b_micro.Bres{
		ResA: res.ResA,
		ResB: res.ResB,
	}, nil
}

func NewBClient(connStr string) (b_micro.Client, error) {
	conn, err := grpc.Dial(connStr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &bClient{
		rpcConn:     conn,
		bClientConn: pb.NewBMicroserviceClient(conn),
	}, nil
}
