package server

import (
	"context"
	"errors"
	"fmt"
	"grpc_micro/b_micro/b_protopb"
	"math/rand"
)

type BServer struct {
}

func (s *BServer) GetB(ctx context.Context, b *b_protopb.BRequest) (*b_protopb.BResponse, error) {
	if b.ParamA == 0 {
		return nil, errors.New("param A cannot be null")
	}
	return &b_protopb.BResponse{
		ResA: fmt.Sprintf("ResA %v", b.ParamA),
		ResB: rand.Int63(),
	}, nil
}
