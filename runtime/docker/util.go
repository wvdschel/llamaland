package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

func DumpInfo() error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	info, err := apiClient.Info(ctx)
	if err != nil {
		return fmt.Errorf("failed to get docker info: %w", err)
	}

	dump, err := json.MarshalIndent(info, "", "  ")

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "%s\n", dump)
	return nil
}
