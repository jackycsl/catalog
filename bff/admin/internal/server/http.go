package server

import (
	jwt2 "github.com/golang-jwt/jwt/v4"
	v1 "github.com/jackycsl/catalog/api/bff/admin/v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jackycsl/catalog/bff/admin/internal/conf"
	"github.com/jackycsl/catalog/bff/admin/internal/service"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, s *service.ShopAdmin, tp *tracesdk.TracerProvider, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(MetricSeconds)),
				metrics.WithRequests(prom.NewCounter(MetricRequests)),
			),
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.ApiKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
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
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", h)
	srv.Handle("/metrics", promhttp.Handler())
	v1.RegisterShopAdminHTTPServer(srv, s)
	return srv
}
