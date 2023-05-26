//go:build wireinject
// +build wireinject

package app

import (
	"go-template-wire/configs"
	"go-template-wire/internal/controller"
	"go-template-wire/internal/middleware"
	"go-template-wire/internal/module"
	"go-template-wire/internal/router"
	"go-template-wire/internal/server"
	"go-template-wire/pkg/databases"
	httpclient "go-template-wire/pkg/http_client"
	"go-template-wire/pkg/logger"
	"go-template-wire/pkg/tracing"

	"github.com/google/wire"
)

var infraSet = wire.NewSet(
	logger.New,
	databases.DatabaseSet,
	tracing.Init,
	httpclient.New,
)

// A Wire injector function that initialize all the app's dependencies
// The return will be filled in by Wire with providers from the provider sets in wire.Build
func initDeps(cfg *configs.Config) (*server.Server, error) {
	wire.Build(
		infraSet,
		middleware.AuthMiddlewareSet,
		module.ModuleSet,
		controller.ControllerSet,
		router.NewEngine,
		server.New,
	)
	return &server.Server{}, nil
}
