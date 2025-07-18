package common

import (
	"context"
	"io"

	"github.com/wvdschel/llamaland/cmd/maestrod/config"
)

type Service interface {
	Prepare(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Logs(ctx context.Context) (io.ReadCloser, io.ReadCloser, error)
}

type Runtime interface {
	NewService(service *config.Service) Service
}
