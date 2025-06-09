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
	ModelData ModelData          `json:"model_data,omitempty"`
	Services  map[string]Service `json:"services,omitempty"`

	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
	TLS      TLS    `json:"tls,omitempty"`

	filename string
}

type ServiceType string

type Service struct {
	Type           ServiceType    `json:"type,omitempty"`
	Spec           map[string]any `json:"spec"`
	RequestLogging RequestLogging `json:"request_logging,omitempty"`
	Models         []string       `json:"models,omitempty"`
}

type ModelData struct {
	Location string `json:"location,omitempty"`
}

type RequestLogging struct {
	Enabled bool `json:"enabled,omitempty"`
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
		ModelData: ModelData{
			Location: xdg.DataHome() + DIRNAME,
		},
		Services: map[string]Service{
			"/deepseek-r1-qwen3-8b": {
				Type: "container",
				Spec: map[string]any{
					"image": "llamaherd/llama-cpp:latest",
				},
				RequestLogging: RequestLogging{
					Enabled: true,
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
