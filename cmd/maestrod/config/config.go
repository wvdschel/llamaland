package config

import (
	"encoding/json"
	"os"
)

type Config struct{}

func LoadFromFile(filename string) (*Config, error) {
	c := &Config{}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
