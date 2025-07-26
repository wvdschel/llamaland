package monitor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wvdschel/llamaland/llamaland/runtime/docker"
)

func TestGPUList(t *testing.T) {
	rt, err := docker.NewRuntime(docker.DefaultOpts().AutodetectRuntime())
	require.NoError(t, err)

	gpus, err := ListGPUs(context.Background(), rt)
	require.NoError(t, err)

	assert.GreaterOrEqual(t, 1, len(gpus))
}
