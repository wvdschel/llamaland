package amd

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/llamaland/runtime/docker"
	"github.com/wvdschel/llamaland/testdata"
)

func TestRcomSmiParser(t *testing.T) {
	info, err := ParseRocmInfo(testdata.RocmSmiJSON)
	require.NoError(t, err)

	require.Equal(t, 1, len(info.GPUs))
	gpu := info.GPUs["card0"]
	assert.Equal(t, 668, gpu.CoreClockMhz)
	assert.Equal(t, 1, gpu.CoreUsePercent)
	assert.Equal(t, int64(536870912), gpu.MemoryTotal)
	assert.Equal(t, int64(415236096), gpu.MemoryUsed)
	assert.Equal(t, 7.006, gpu.PowerDrawW)
	assert.Equal(t, "Radeon 8060S Graphics", gpu.Name)
	assert.Equal(t, 43.0, gpu.TemperatureC)
}

func Test_RocmInfo(t *testing.T) {
	rt, err := docker.NewRuntime(docker.DefaultOpts().AutodetectRuntime())
	require.NoError(t, err)

	svc := rt.NewService(&config.Service{
		Type: "docker",
		Spec: map[string]any{
			"image": "quay.io/llamaland/rcom-smi:latest",
		},
	})

	require.NoError(t, svc.Prepare(context.Background()))
	require.NoError(t, svc.Start(context.Background()))

	stdout, stderr, err := svc.Logs(context.Background())
	require.NoError(t, err)

	go func() { _, _ = io.ReadAll(stderr) }()
	logs, err := io.ReadAll(stdout)
	assert.NoError(t, err)
	deviceInfo, err := ParseRocmInfo(logs)
	require.NoError(t, err)
	assert.Greater(t, len(deviceInfo.GPUs), 0)
}
