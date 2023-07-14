package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/gorilla/handlers"
	v1 "mall-shop-be/api/helloworld/v1"
	"mall-shop-be/internal/conf"
	"mall-shop-be/internal/pkg/auth"
	"mall-shop-be/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, Mall *service.MallService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.AUTH()).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
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
	v1.RegisterMallHTTPServer(srv, Mall)
	return srv
}
func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/api.helloworld.v1.Mall/Login":    {},
		"/api.helloworld.v1.Mall/Register": {},
	}

	return func(ctx context.Context, operation string) bool {
		log.Info(operation)
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}
