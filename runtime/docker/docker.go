package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/wvdschel/llamaland/cmd/maestrod/config"
)

type Runtime struct {
	dockerClient      *client.Client
	dockerRuntimeName string
}

func NewRuntime(opts *opts) (*Runtime, error) {
	// TODO check available runtimes, pick one:
	// - nvidia
	// - rocm?

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	info, err := apiClient.Info(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get docker info: %w", err)
	}

	// TODO: process info.DiscoveredDevices
	_ = info

	return &Runtime{
		dockerClient:      apiClient,
		dockerRuntimeName: opts.runtime,
	}, nil
}

func (r *Runtime) NewService(service *config.Service) *Service {
	return &Service{
		runtime: r,
		cfg:     service,
	}
}
