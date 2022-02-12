//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jackycsl/catalog/bff/interface/internal/biz"
	"github.com/jackycsl/catalog/bff/interface/internal/conf"
	"github.com/jackycsl/catalog/bff/interface/internal/data"
	"github.com/jackycsl/catalog/bff/interface/internal/server"
	"github.com/jackycsl/catalog/bff/interface/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
