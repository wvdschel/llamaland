package opencl

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/llamaland/runtime/docker"
)

func Test_Clinfo(t *testing.T) {
	rt, err := docker.NewRuntime(docker.DefaultOpts().AutodetectRuntime())
	require.NoError(t, err)

	svc := rt.NewService(&config.Service{
		Type: "docker",
		Spec: map[string]any{
			"image": "quay.io/llamaland/clinfo:latest",
		},
	})

	require.NoError(t, svc.Prepare(context.Background()))
	require.NoError(t, svc.Start(context.Background()))

	stdout, stderr, err := svc.Logs(context.Background())
	require.NoError(t, err)

	go func() { _, _ = io.ReadAll(stderr) }()
	logs, err := io.ReadAll(stdout)
	assert.NoError(t, err)
	deviceInfo, err := ParseCLInfoOutput(string(logs))
	require.NoError(t, err)
	assert.Greater(t, len(deviceInfo), 0)
}
