package service

import (
	"context"
	"encoding/hex"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	v1 "helloworld/api/helloworld/v1"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	reply := &v1.HelloReply{}
	a, err := anypb.New(&v1.Data{
		FieldA: "A",
		FieldB: "B",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(a)

	reply.Any = a
	marshal, err := proto.Marshal(reply)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(marshal))
	t.Log(hex.EncodeToString(marshal))
}

func TestClient(t *testing.T) {
	newClient, err2 := http.NewClient(context.Background(),
		http.WithEndpoint("localhost:8000"),
		http.WithMiddleware(),
	)
	if err2 != nil {
		t.Error(err2)
		return
	}
	client := v1.NewGreeterHTTPClient(newClient)

	rsp, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "hubo"})
	if err != nil {
		t.Error(err)
	}

	t.Log(rsp)
}

func TestGrpcClient(t *testing.T) {
	newClient, err2 := grpc.Dial(context.Background(),
		grpc.WithEndpoint("localhost:9000"),
		grpc.WithMiddleware(
			clientRespMiddleware,
		),
		grpc.WithOptions(
			grpc2.WithTransportCredentials(insecure.NewCredentials()),
			grpc2.WithUnaryInterceptor(UnaryInterceptor()),
		),
		grpc.WithTimeout(time.Hour),
	)
	insecure.NewCredentials()
	if err2 != nil {
		t.Error(err2)
		return
	}
	client := v1.NewGreeterClient(newClient)

	rsp, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "hubo"})
	if err != nil {
		t.Error(err)
	}

	t.Log(rsp)
}

func clientRespMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		return
	}
}

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

		return errors.New(int(realReply.Code), "", realReply.Msg)
	}
}
