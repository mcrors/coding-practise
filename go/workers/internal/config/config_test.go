package config

import (
	"fmt"
	"os"
	"testing"
)

// test that defaults are applied if no file is given and no env var is set
func TestLoad_Defaults(t *testing.T) {
	clearEnv(t)

	cfg, err := Load("")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	d := defaults()
	if cfg.NumWorkers != d.NumWorkers {
		t.Errorf("NumWorkers: got %d, want %d", cfg.NumWorkers, d.NumWorkers)
	}
}

// test that yaml file overrides defaults
func TestLoad_FileOverridesDefaults(t *testing.T) {
	clearEnv(t)

	f := writeTempConfig(t, `
num_workers: 5
`)
	cfg, err := Load(f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.NumWorkers != 5 {
		t.Errorf("NumWorkers: got %d, want %d", cfg.NumWorkers, 5)
	}
}

// test that an error is thrown for incorrect data types in the yaml file
func TestLoad_InvalidYamlValues(t *testing.T) {
	clearEnv(t)

	f := writeTempConfig(t, `
num_workers: true
`)

	_, err := Load(f)

	if err == nil {
		t.Error("expected err for invalid YAML values")
	}
}

// test invalid env var
func TestLoad_InvalidEnv(t *testing.T) {
	clearEnv(t)

	t.Setenv("WORKERS_NUM_WORKERS", "notanumber")
	_, err := Load("")
	if err == nil {
		t.Error("expected error for invalid WORKERS_NUM_WORKERS")
	}
}

// test that env var overrides YAML
func TestLoad_EnvVarOverridesYamlFile(t *testing.T) {
	clearEnv(t)

	f := writeTempConfig(t, `
num_workers: 5
`)

	expected := 3
	t.Setenv("WORKERS_NUM_WORKERS", fmt.Sprintf("%d", expected))

	cfg, err := Load(f)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if cfg.NumWorkers != expected {
		t.Errorf("NumWorkers: got %d, want %d", cfg.NumWorkers, expected)
	}
}

// test that env var overrides defaults
func TestLoad_EnvVarOverridesDefaults(t *testing.T) {
	clearEnv(t)

	expected := 3
	t.Setenv("WORKERS_NUM_WORKERS", fmt.Sprintf("%d", expected))

	cfg, err := Load("")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if cfg.NumWorkers != expected {
		t.Errorf("NumWorkers: got %d, want %d", cfg.NumWorkers, expected)
	}
}

// helpers
func clearEnv(t *testing.T) {
	t.Helper()
	varNames := []string{
		"WORKERS_CONFIG_PATH", "WORKERS_NUM_WORKERS",
	}

	for _, v := range varNames {
		t.Setenv(v, "")
		os.Unsetenv(v)
	}
}

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "config.yaml")
	if err != nil {
		t.Fatalf("creating temp config: %v", err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("writing temp config: %v", err)
	}
	return f.Name()
}
