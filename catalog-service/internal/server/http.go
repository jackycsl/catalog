package server

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/jackycsl/catalog/api/catalog/service/v1"
	"github.com/jackycsl/catalog/catalog-service/internal/conf"
	"github.com/jackycsl/catalog/catalog-service/internal/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, catalog *service.CatalogService, tp *tracesdk.TracerProvider, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(MetricSeconds)),
				metrics.WithRequests(prom.NewCounter(MetricRequests)),
			),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	v1.RegisterCatalogHTTPServer(srv, catalog)
	return srv
}
