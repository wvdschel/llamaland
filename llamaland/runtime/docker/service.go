package docker

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/google/uuid"
	"github.com/wvdschel/llamaland/cmd/maestrod/config"
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
	img, ok := s.cfg.Spec["image"]
	if !ok {
		return errors.New("spec.image is required for docker services")
	}

	s.imageName, ok = img.(string)
	if !ok {
		return errors.New("spec.image must be a string")
	}
	return nil
}

func (s *Service) Prepare(ctx context.Context) error {
	if err := s.loadSpec(); err != nil {
		return err
	}

	imageList, err := s.docker().ImageList(ctx, image.ListOptions{
		Filters: filters.NewArgs(
			filters.KeyValuePair{
				Key:   "reference",
				Value: s.imageName,
			},
		),
	})
	if err != nil {
		fmt.Printf("failed to list docker images: %v\n", err)
		return err
	}

	if len(imageList) == 0 {
		pullLog, err := s.docker().ImagePull(ctx, s.imageName, image.PullOptions{})
		if err != nil {
			fmt.Printf("failed to pull docker image: %v\n", err)
			return err
		}

		scanner := bufio.NewScanner(pullLog)
		for scanner.Scan() {
		}
		pullLog.Close()
	}

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
			Runtime: s.runtime.opts.runtime,
			Resources: container.Resources{ // TODO derive from options
				DeviceRequests: s.runtime.opts.deviceRequests,
				Devices:        s.runtime.opts.devices,
			},
			Privileged: true,
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

func (s *Service) Logs(ctx context.Context) (io.ReadCloser, io.ReadCloser, error) {
	stream, err := s.docker().ContainerLogs(ctx, s.containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		return nil, nil, err
	}

	stdout_r, stdout := io.Pipe()
	stderr_r, stderr := io.Pipe()

	go func() {
		if _, err := stdcopy.StdCopy(stdout, stderr, stream); err != nil {
			_ = stderr.CloseWithError(err)
			_ = stdout.CloseWithError(err)
		} else {
			_ = stdout.Close()
			_ = stderr.Close()
		}
	}()

	return stdout_r, stderr_r, nil
}
