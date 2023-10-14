package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			httpLog(logger),
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
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}

func httpLog(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code   int32
				reason string
				//kind      string
				//operation string
				level log.Level
				stack string
			)
			startTime := time.Now()
			time.Sleep(time.Millisecond)
			request, _ := http.RequestFromServerContext(ctx)

			reply, err = handler(ctx, req)

			kvs := []any{
				"path", request.RequestURI,
				"args", req,
			}

			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
				level, stack = extractError(err)
				kvs = append(kvs, []any{
					"code", code,
					"reason", reason,
					"stack", stack,
				}...)
			}

			kvs = append(kvs, []any{
				"latency", time.Since(startTime),
			}...)

			_ = log.WithContext(ctx, logger).Log(level, kvs...)
			return
		}
	}
}

func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}
