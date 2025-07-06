package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/wvdschel/llamaland/xdg"
)

const DIRNAME = "llamaland"

var DefaultFilename = xdg.ConfigHome() + "/" + DIRNAME + "/config.json"

type Config struct {
	Models   map[string]Model   `json:"models,omitempty"`
	Services map[string]Service `json:"services,omitempty"`
	Runtimes map[string]Runtime `json:"runtimes,omitempty"`

	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
	TLS      TLS    `json:"tls,omitempty"`

	StorageLocations StorageConfig `json:"storage_locations,omitempty"`

	filename string
}

type ServiceType string

type Service struct {
	Type    ServiceType    `json:"type,omitempty"`
	Spec    map[string]any `json:"spec"`
	Options ServiceOptions `json:"service_options,omitempty"`
	Models  []string       `json:"models,omitempty"`
}

type StorageConfig struct {
	Models string `json:"models,omitempty"`
}

type Model struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Size        int    `json:"size,omitempty"`
	Description string `json:"description,omitempty"`
}

type Runtime struct{}

type ServiceOptions struct {
	RequestLogging bool `json:"request_logging,omitempty"`
}

type TLS struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

func LoadFromFile(filename string) (*Config, error) {
	c := Default()
	c.filename = filename

	f, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) && filename == DefaultFilename {
		err = os.MkdirAll(xdg.ConfigHome()+"/"+DIRNAME, 0755)
		if err != nil {
			return nil, err
		}
		return c, c.Save()
	}
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

func (c *Config) Save() error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(c.filename, b, 0644)
}

func Default() *Config {
	return &Config{
		StorageLocations: StorageConfig{
			Models: xdg.DataHome() + DIRNAME + "/models",
		},
		Services: map[string]Service{
			"/deepseek-r1-qwen3-8b": {
				Type: "container",
				Spec: map[string]any{
					"image": "llamaland/llama-cpp:latest",
				},
				Options: ServiceOptions{
					RequestLogging: true,
				},
				Models: []string{
					"https://huggingface.co/unsloth/DeepSeek-R1-0528-Qwen3-8B-GGUF/DeepSeek-R1-0528-Qwen3-8B-UD-Q4_K_XL.gguf",
				},
			},
		},
		Hostname: "0.0.0.0",
		Port:     18080,
		TLS: TLS{
			Cert: "",
			Key:  "",
		},
	}
}
