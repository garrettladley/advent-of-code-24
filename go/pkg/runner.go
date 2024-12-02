package pkg

import (
	"context"
	"io"
)

type Runner interface {
	Run(ctx context.Context, r io.Reader) (string, error)
}
