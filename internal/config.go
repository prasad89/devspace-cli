package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// getConfig initializes and loads the config file
func GetConfig() (*ini.File, string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	configPath := filepath.Join(homeDir, ".devspace", "config.ini")

	if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
		fmt.Printf("Failed to create config directory: %v\n", err)
		os.Exit(1)
	}

	cfg, err := ini.Load(configPath)
	if err != nil {
		cfg = ini.Empty()

		if err := cfg.SaveTo(configPath); err != nil {
			fmt.Printf("Failed to save config file: %v\n", err)
			os.Exit(1)
		}
	}

	return cfg, configPath
}
