package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
		//grpc.UnaryInterceptor(UnaryInterceptor()),
	}
	opts = append(opts, buildGrpcOption(c)...)

	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}

func buildGrpcOption(c *conf.Server) []grpc.ServerOption {
	var opts []grpc.ServerOption
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	return opts
}
