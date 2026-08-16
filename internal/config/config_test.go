package config_test

import (
	"os"
	"testing"

	"github.com/apple-wallet-automation/internal/config"
)

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "config-*.yml")
	if err != nil {
		t.Fatalf("failed to create temp config: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}
	f.Close()
	return f.Name()
}

func TestLoad_ValidConfig(t *testing.T) {
	raw := `
users:
  - username: alice
    api_key: key-alice
  - username: bob
    api_key: key-bob

categories:
  Groceries:
    - yerevan city
    - carrefour
  Restaurant:
    - kfc
    - starbucks

storage:
  data_dir: ./data
`
	path := writeTempConfig(t, raw)
	cfg, err := config.Load(path)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if len(cfg.Users) != 2 {
		t.Errorf("expected 2 users, got %d", len(cfg.Users))
	}
	if cfg.Users[0].Username != "alice" {
		t.Errorf("expected alice, got %s", cfg.Users[0].Username)
	}
	if cfg.Users[1].APIKey != "key-bob" {
		t.Errorf("expected key-bob, got %s", cfg.Users[1].APIKey)
	}
	if cfg.Storage.DataDir != "./data" {
		t.Errorf("expected ./data, got %s", cfg.Storage.DataDir)
	}
	if len(cfg.Categories["Groceries"]) != 2 {
		t.Errorf("expected 2 grocery keywords, got %d", len(cfg.Categories["Groceries"]))
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := config.Load("/nonexistent/path/config.yml")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	path := writeTempConfig(t, "key: [unclosed bracket")
	_, err := config.Load(path)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

func TestFindUserByAPIKey_Found(t *testing.T) {
	raw := `
users:
  - username: alice
    api_key: key-alice
storage:
  data_dir: ./data
`
	cfg, _ := config.Load(writeTempConfig(t, raw))

	user, ok := cfg.FindUserByAPIKey("key-alice")
	if !ok {
		t.Fatal("expected to find user, got false")
	}
	if user.Username != "alice" {
		t.Errorf("expected alice, got %s", user.Username)
	}
}

func TestFindUserByAPIKey_NotFound(t *testing.T) {
	raw := `
users:
  - username: alice
    api_key: key-alice
storage:
  data_dir: ./data
`
	cfg, _ := config.Load(writeTempConfig(t, raw))

	_, ok := cfg.FindUserByAPIKey("wrong-key")
	if ok {
		t.Error("expected false for unknown api key, got true")
	}
}

func TestFindUserByAPIKey_EmptyKey(t *testing.T) {
	raw := `
users:
  - username: alice
    api_key: key-alice
storage:
  data_dir: ./data
`
	cfg, _ := config.Load(writeTempConfig(t, raw))

	_, ok := cfg.FindUserByAPIKey("")
	if ok {
		t.Error("expected false for empty api key, got true")
	}
}
