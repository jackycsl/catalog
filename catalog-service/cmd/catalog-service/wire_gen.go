// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/catalog-service/internal/conf"
	"github.com/jackycsl/catalog/catalog-service/internal/data"
	"github.com/jackycsl/catalog/catalog-service/internal/server"
	"github.com/jackycsl/catalog/catalog-service/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	redisClient := data.NewRedisClient(confData, logger)
	sender := data.NewKafkaProducer(confData, logger)
	dataData, cleanup, err := data.NewData(client, redisClient, sender, logger)
	if err != nil {
		return nil, nil, err
	}
	gameRepo := data.NewGameRepo(dataData, logger)
	gameUseCase := biz.NewGameUseCase(gameRepo, logger)
	catalogService := service.NewCatalogService(gameUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, catalogService, tracerProvider, logger)
	grpcServer := server.NewGRPCServer(confServer, catalogService, tracerProvider, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
