package b_micro

import (
	"context"
)

type Client interface {
	GetB(ctx context.Context, param1 int64) (*Bres, error)
}
