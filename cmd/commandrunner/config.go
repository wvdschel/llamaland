package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	Commands map[string]CommandConfig `json:"commands"`
}

type CommandConfig struct {
	Command string            `json:"command"`
	Args    []string          `json:"args,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
	Timeout time.Duration     `json:"timeout,omitempty"`
}

func loadConfig() (*Config, error) {
	configData, err := os.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return nil, err
	}

	if config.Port == 0 {
		return nil, errors.New("error loading config: port is required")
	}

	if config.Host == "" {
		return nil, errors.New("error loading config: host is required")
	}

	if len(config.Commands) == 0 {
		return nil, errors.New("error loading config: at least one command is required")
	}

	return &config, nil
}

func (c *CommandConfig) Environment() []string {
	env := []string{}
	for k, v := range c.Env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}
