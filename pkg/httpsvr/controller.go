package httpsvr

import (
	"context"

	"github.com/dup2X/gomicrosvr/pkg/idl"
)

type Controller interface {
	Bind()
	Do(ctx context.Context, input idl.Request) (idl.Response, error)
}
