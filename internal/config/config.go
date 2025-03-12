package config

import (
	"flag"

	"github.com/joho/godotenv"
)

var configPath string

func initPath() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func Load() error {
	initPath()
	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}
