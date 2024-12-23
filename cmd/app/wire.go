//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"go-gin-framework/internal/base/conf"
	"go-gin-framework/internal/base/data"
	"go-gin-framework/internal/controller"
	"go-gin-framework/internal/repository"
	"go-gin-framework/internal/router"
	"go-gin-framework/internal/server"
	"go-gin-framework/internal/service"

	"github.com/google/wire"
)

func initializeServer(debug bool, dataConf *conf.Database) (*server.Server, error) {
	panic(wire.Build(
		data.NewDB,
		data.NewData,
		repository.ProviderSetRepository,
		service.ProviderSetService,
		controller.ProviderSetController,
		router.ProviderSetRouter,
		server.ProviderSetServer,
	))
}
