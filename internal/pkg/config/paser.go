package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Users  []User `yaml:"users"`
	APIKey string `yaml:"apikey"`
}

type User struct {
	Name string `yaml:"name"`
	Pass string `yaml:"pass"`
	Role string `yaml:"role"`
}

func ParseConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
