package server

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/pkg/errs"
	"helloworld/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			//respMiddleware,
		),
		grpc.UnaryInterceptor(UnaryInterceptor()),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}

func respMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		a, err := anypb.New(reply.(proto.Message))
		if err != nil {
			log.Error("响应必须是protobuf类型")
			panic(err)
		}

		var code int
		var msg string

		var iError errs.IError
		if ok := errors.As(err, &iError); ok {
			code = iError.Code()
			msg = iError.Error()
		}

		reply = &v1.Reply{
			Data: a,
			Code: int32(code),
			Msg:  msg,
		}
		return
	}
}

func UnaryInterceptor() grpc2.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc2.UnaryServerInfo, handler grpc2.UnaryHandler) (resp interface{}, err error) {
		log.Info(req)

		resp, err = handler(ctx, req)

		var code int
		var msg string

		var iError errs.IError
		if ok := errors.As(err, &iError); ok {
			code = iError.Code()
			msg = iError.Error()
		}

		a, err := anypb.New(resp.(proto.Message))
		if err != nil {
			log.Error(err)
		}
		resp = &v1.Reply{
			Data: a,
			Code: int32(code),
			Msg:  msg,
		}

		return resp, nil
	}
}
