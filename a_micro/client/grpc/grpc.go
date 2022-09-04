package grpc

import (
	"context"
	"google.golang.org/grpc"
	"grpc_micro/a_micro"
	pb "grpc_micro/a_micro/a_protopb"
)

var _ a_micro.Client = (*aClient)(nil)

type aClient struct {
	rpcConn     *grpc.ClientConn
	aClientConn pb.AMicroserviceClient
}

func (c *aClient) GetA(ctx context.Context, param1 string, param2 int64) (*a_micro.Ares, error) {
	res, err := c.aClientConn.GetA(ctx, &pb.ARequest{
		ParamA: param1,
		ParamB: param2,
	})
	if err != nil {
		return nil, err
	}

	return &a_micro.Ares{
		ResA: res.ResA,
		ResB: res.ResB,
	}, nil
}

func NewAClient(connStr string) (a_micro.Client, error) {
	conn, err := grpc.Dial(connStr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &aClient{
		rpcConn:     conn,
		aClientConn: pb.NewAMicroserviceClient(conn),
	}, nil
}
