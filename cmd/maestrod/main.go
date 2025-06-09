package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/runtime/docker"
)

var configFilename = flag.String("config", config.DefaultFilename, "config file")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadFromFile(*configFilename)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// TODO automatically select runtime / support multiple different runtimes at once

	rt, err := docker.NewRuntime(docker.DefaultOpts().WithRuntime("nvidia"))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	svc := rt.NewService(&config.Service{
		Type: "container",
		Spec: map[string]any{
			"image": "llamaherd/nvidia-smi:latest",
		},
	})

	if err := svc.Prepare(ctx); err != nil {
		log.Fatalf("failed to start container: %v", err)
	}

	if err := svc.Start(ctx); err != nil {
		log.Fatalf("failed to start container: %v", err)
	}

	time.Sleep(3 * time.Second)

	logReader, _, err := svc.Logs(ctx)
	if err != nil {
		log.Fatalf("failed to get container logs: %v", err)
	}
	logs, err := io.ReadAll(logReader)
	if err != nil {
		log.Fatalf("failed to read container logs: %v", err)
	}

	// TODO check host config from running container while starting an nvidia container interactively from the shell
	// to see what's wrong / missing
	fmt.Println(string(logs))

	_, _ = cfg, rt
}
