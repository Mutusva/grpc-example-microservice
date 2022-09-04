package a_micro

import (
	"context"
)

type Client interface {
	GetA(ctx context.Context, param1 string, param2 int64) (*Ares, error)
}
