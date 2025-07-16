package docker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/llamaland/runtime/common"
)

type Runtime struct {
	dockerClient *client.Client
	opts         opts
}

func AutodetectRuntime() (common.Runtime, error) {
	return nil, errors.New("not implemented")
}

func NewRuntime(opts *opts) (common.Runtime, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	if opts == nil {
		opts = DefaultOpts()
	}

	return &Runtime{
		dockerClient: apiClient,
		opts:         *opts,
	}, nil
}

func (r *Runtime) NewService(service *config.Service) common.Service {
	return &Service{
		runtime: r,
		cfg:     service,
	}
}

func SystemInfo() (system.Info, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return system.Info{}, err
	}
	defer apiClient.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return apiClient.Info(ctx)
}
