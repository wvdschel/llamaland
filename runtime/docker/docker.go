package docker

import (
	"github.com/docker/docker/client"
	"github.com/wvdschel/compute-maestro/cmd/maestrod/config"
)

type Runtime struct {
	dockerClient      *client.Client
	dockerRuntimeName string
}

func NewRuntime(dockerRuntimeName string) (*Runtime, error) {
	// TODO check available runtimes, pick one:
	// - nvidia
	// - rocm?

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &Runtime{
		dockerClient:      apiClient,
		dockerRuntimeName: dockerRuntimeName,
	}, nil
}

func (r *Runtime) NewService(service *config.Service) *Service {
	return &Service{
		runtime: r,
		cfg:     service,
	}
}
