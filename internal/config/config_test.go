package config_test

import (
	"os"
	"testing"

	"github.com/apple-wallet-automation/internal/config"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "*.yml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	f.Close()
	return f.Name()
}

func validConfigYML() string {
	return `
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
}

func validCredentialsYML() string {
	return `
users:
  - username: alice
    api_key: key-alice
  - username: bob
    api_key: key-bob
`
}

func TestLoad_ValidConfig(t *testing.T) {
	cfgPath := writeTempFile(t, validConfigYML())
	credPath := writeTempFile(t, validCredentialsYML())

	cfg, err := config.Load(cfgPath, credPath)
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

func TestLoad_ConfigFileNotFound(t *testing.T) {
	credPath := writeTempFile(t, validCredentialsYML())
	_, err := config.Load("/nonexistent/config.yml", credPath)
	if err == nil {
		t.Error("expected error for missing config file, got nil")
	}
}

func TestLoad_CredentialsFileNotFound(t *testing.T) {
	cfgPath := writeTempFile(t, validConfigYML())
	_, err := config.Load(cfgPath, "/nonexistent/credentials.yml")
	if err == nil {
		t.Error("expected error for missing credentials file, got nil")
	}
}

func TestLoad_InvalidConfigYAML(t *testing.T) {
	cfgPath := writeTempFile(t, "key: [unclosed bracket")
	credPath := writeTempFile(t, validCredentialsYML())
	_, err := config.Load(cfgPath, credPath)
	if err == nil {
		t.Error("expected error for invalid config YAML, got nil")
	}
}

func TestLoad_InvalidCredentialsYAML(t *testing.T) {
	cfgPath := writeTempFile(t, validConfigYML())
	credPath := writeTempFile(t, "key: [unclosed bracket")
	_, err := config.Load(cfgPath, credPath)
	if err == nil {
		t.Error("expected error for invalid credentials YAML, got nil")
	}
}

func TestFindUserByAPIKey_Found(t *testing.T) {
	cfg, _ := config.Load(
		writeTempFile(t, validConfigYML()),
		writeTempFile(t, validCredentialsYML()),
	)

	user, ok := cfg.FindUserByAPIKey("key-alice")
	if !ok {
		t.Fatal("expected to find user, got false")
	}
	if user.Username != "alice" {
		t.Errorf("expected alice, got %s", user.Username)
	}
}

func TestFindUserByAPIKey_NotFound(t *testing.T) {
	cfg, _ := config.Load(
		writeTempFile(t, validConfigYML()),
		writeTempFile(t, validCredentialsYML()),
	)

	_, ok := cfg.FindUserByAPIKey("wrong-key")
	if ok {
		t.Error("expected false for unknown api key, got true")
	}
}

func TestFindUserByAPIKey_EmptyKey(t *testing.T) {
	cfg, _ := config.Load(
		writeTempFile(t, validConfigYML()),
		writeTempFile(t, validCredentialsYML()),
	)

	_, ok := cfg.FindUserByAPIKey("")
	if ok {
		t.Error("expected false for empty api key, got true")
	}
}
