// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackycsl/catalog/user-service/internal/biz"
	"github.com/jackycsl/catalog/user-service/internal/conf"
	"github.com/jackycsl/catalog/user-service/internal/data"
	"github.com/jackycsl/catalog/user-service/internal/server"
	"github.com/jackycsl/catalog/user-service/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, tracerProvider, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, tracerProvider, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}