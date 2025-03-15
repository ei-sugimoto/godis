package env

import (
	"os"

	"github.com/ei-sugimoto/godis/internal/pkg/err"
)

func GetPort() (string, error) {
	port := os.Getenv("GODIS_PORT")

	if port == "" {
		return "", err.ErrPortEmpty
	}

	return port, nil
}

func GetConfigPath() (string, error) {
	configPath := os.Getenv("GODIS_CONFIG_PATH")

	if configPath == "" {
		return "", err.ErrConfigPathEmpty
	}

	return configPath, nil
}
