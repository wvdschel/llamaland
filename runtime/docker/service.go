package docker

import (
	"bufio"
	"context"
	"errors"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/wvdschel/compute-maestro/cmd/maestrod/config"
)

type Service struct {
	runtime *Runtime
	cfg     *config.Service

	imageName   string
	containerID string
}

func (s *Service) docker() *client.Client {
	return s.runtime.dockerClient
}

func (s *Service) loadSpec() error {
	img, ok := s.cfg.Spec["image_name"]
	if !ok {
		return errors.New("spec.image_name is required for docker services")
	}

	s.imageName, ok = img.(string)
	if !ok {
		return errors.New("spec.image_name must be a string")
	}
	return nil
}

func (s *Service) Prepare(ctx context.Context) error {
	if err := s.loadSpec(); err != nil {
		return err
	}

	pullLog, err := s.docker().ImagePull(ctx, s.imageName, image.PullOptions{})
	if err != nil {
		fmt.Printf("failed to pull docker image: %v\n", err)
		return err
	}

	scanner := bufio.NewScanner(pullLog)
	for scanner.Scan() {
	}
	pullLog.Close()

	createResp, err := s.docker().ContainerCreate(ctx,
		&container.Config{
			Image: s.imageName,
		},
		&container.HostConfig{
			DNS: []string{
				"8.8.8.8",
				"8.8.4.4",
				"2001:4860:4860::8888",
				"2001:4860:4860::8844"},
			DNSSearch: []string{
				"8.8.8.8",
				"8.8.4.4",
				"2001:4860:4860::8888",
				"2001:4860:4860::8844"},
			PublishAllPorts: false,
			NetworkMode:     network.NetworkDefault,
			RestartPolicy: container.RestartPolicy{
				Name:              container.RestartPolicyDisabled,
				MaximumRetryCount: 0,
			},
			Runtime: s.runtime.dockerRuntimeName,
		},
		nil, nil, uuid.New().String())
	if err != nil {
		fmt.Printf("failed to create container: %v\n", err)
		return err
	}

	s.containerID = createResp.ID

	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return s.docker().ContainerStart(ctx, s.containerID, container.StartOptions{})
}

func ptr[T any](v T) *T {
	return &v
}

func (s *Service) Stop(ctx context.Context) error {
	err := s.docker().ContainerStop(ctx, s.containerID, container.StopOptions{
		Timeout: ptr(5),
	})

	if err != nil {
		return err
	}

	err = s.docker().ContainerRemove(ctx, s.containerID, container.RemoveOptions{
		Force: true,
	})

	return err
}
