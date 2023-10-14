package test

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	v1 "helloworld/api/helloworld/v1"
	"log"
)

func UnaryInterceptor() grpc2.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc2.ClientConn, invoker grpc2.UnaryInvoker, opts ...grpc2.CallOption) error {
		//log.Print(method)
		//log.Print(req)

		realReply := &v1.Reply{}

		err := invoker(ctx, method, req, realReply, cc, opts...)

		if err != nil {
			log.Print(err)
			return err
		}

		err = anypb.UnmarshalTo(realReply.Data, reply.(proto.Message), proto.UnmarshalOptions{})
		if err != nil {
			log.Print(err)
			return err
		}

		if realReply.Code == 0 || realReply.Code == 200 {
			return nil
		}

		return nil
	}
}

func clientRespMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		return
	}
}
