package main

import (
	"go-gin-framework/internal/base/conf"
	"path/filepath"
)

const (
	DefaultConfigPath = "/configs/config.yaml"
)

func main() {
	Execute()
}

func RunApp() {
	config, err := conf.LoadConfig(filepath.Join(".", DefaultConfigPath))
	if err != nil {
		panic(err)
	}

	server, err := initializeServer(config.Debug, &config.Data.Database, &config.Server)
	if err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}
}
