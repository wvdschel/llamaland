package vulkan

import (
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/llamaland/runtime/docker"
)

func Test_Vulkaninfo(t *testing.T) {
	rt, err := docker.NewRuntime(docker.DefaultOpts())
	require.NoError(t, err)

	svc := rt.NewService(&config.Service{
		Type: "docker",
		Spec: map[string]any{
			"image": "quay.io/llamaland/vulkaninfo:latest",
		},
	})

	require.NoError(t, svc.Prepare(context.Background()))
	require.NoError(t, svc.Start(context.Background()))

	stdout, stderr, err := svc.Logs(context.Background())
	require.NoError(t, stderr.Close())
	require.NoError(t, err)

	logs, err := io.ReadAll(stdout)
	assert.NoError(t, err)
	//fmt.Println(string(logs))

	r := []VulkanDeviceInfo{}
	err = json.Unmarshal(logs, &r)
	require.NoError(t, err)
	assert.Greater(t, len(r), 0)
}
