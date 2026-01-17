package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port                    string `yaml:"port"`
	URLQueueBufferSize      int    `yaml:"url_queue_buffer_size"`
	PlaybackQueueBufferSize int    `yaml:"playback_queue_buffer_size"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &cfg, nil
}
