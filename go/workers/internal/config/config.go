// Package config is used to load the application Config struct
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	NumWorkers int `yaml:"num_workers"`
}

func defaults() Config {
	return Config{
		NumWorkers: 4,
	}
}

// Load reads the config from configPath (YAML), it then applies default and env var
// overrides. If configPath is empty, it defaults to ./config.yaml
// A missing config file is not an Error, it is ignored and the defaults apply
func Load(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = "./config.yaml"
	}
	if v := os.Getenv("WORKERS_CONFIG_PATH"); v != "" {
		configPath = v
	}

	cfg := defaults()

	if err := loadYaml(configPath, &cfg); err != nil {
		return nil, err
	}

	if err := applyEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadYaml(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("parsing config file: %w", err)
	}
	return nil
}

func applyEnv(cfg *Config) error {
	if v := os.Getenv("WORKERS_NUM_WORKERS"); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid WORKERS_NUM_WORKERS %q: must be a positive integer", v)
		}
		cfg.NumWorkers = i
	}
	return nil
}
