package data

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	catalogv1 "github.com/jackycsl/catalog/api/catalog/service/v1"
	"github.com/jackycsl/catalog/bff/interface/internal/conf"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewRegistrar, NewCatalogServiceClient, NewCatalogRepo)

// Data .
type Data struct {
	log *log.Helper
	cc  catalogv1.CatalogClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, cc catalogv1.CatalogClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log: l,
		cc:  cc,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewCatalogServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) catalogv1.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///game.catalog.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return catalogv1.NewCatalogClient(conn)
}
