package main

import (
	"flag"
	"log"

	"github.com/wvdschel/compute-maestro/cmd/maestrod/config"
	"github.com/wvdschel/compute-maestro/xdg"
)

var configFilename = flag.String("config", xdg.ConfigHome()+"/compute-maestro/config.json", "config file")

func main() {
	flag.Parse()

	cfg, err := config.LoadFromFile(*configFilename)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	_ = cfg
}
