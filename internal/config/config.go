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
	Users      []User                `yaml:"users"`
	Categories map[string][]string   `yaml:"categories"`
	Storage    Storage               `yaml:"storage"`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

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
