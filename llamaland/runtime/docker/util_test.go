package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/require"
)

func Test_DumpRunningContainers(t *testing.T) {
	t.SkipNow() // This test exists to dump the running containers to stdout. It is not meant to be run as a test.

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	require.NoError(t, err)

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{})
	require.NoError(t, err)

	b, err := json.MarshalIndent(containers, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(b))
}

func Test_DumpDockerInfo(t *testing.T) {
	t.SkipNow() // This test exists to dump the docker info to stdout. It is not meant to be run as a test.

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	require.NoError(t, err)

	info, err := apiClient.Info(context.Background())
	require.NoError(t, err)

	b, err := json.MarshalIndent(info, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(b))
}
