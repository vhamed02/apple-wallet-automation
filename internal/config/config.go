package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Username string `yaml:"username"`
	APIKey   string `yaml:"api_key"`
}

type Storage struct {
	DataDir string `yaml:"data_dir"`
}

type Config struct {
	Users      []User              `yaml:"users"`
	Categories map[string][]string `yaml:"categories"`
	Storage    Storage             `yaml:"storage"`
}

type credentials struct {
	Users []User `yaml:"users"`
}

func Load(configPath, credentialsPath string) (*Config, error) {
	cf, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer cf.Close()

	var cfg Config
	if err := yaml.NewDecoder(cf).Decode(&cfg); err != nil {
		return nil, err
	}

	cr, err := os.Open(credentialsPath)
	if err != nil {
		return nil, err
	}
	defer cr.Close()

	var creds credentials
	if err := yaml.NewDecoder(cr).Decode(&creds); err != nil {
		return nil, err
	}

	cfg.Users = creds.Users

	return &cfg, nil
}

func (c *Config) FindUserByAPIKey(key string) (User, bool) {
	for _, u := range c.Users {
		if u.APIKey == key {
			return u, true
		}
	}
	return User{}, false
}
