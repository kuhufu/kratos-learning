package test

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "helloworld/api/helloworld/v1"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
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
	newClient, err := grpc.Dial(context.Background(),
		grpc.WithEndpoint("localhost:9000"),
		grpc.WithMiddleware(
		//clientRespMiddleware,
		),
		grpc.WithOptions(
			grpc2.WithTransportCredentials(insecure.NewCredentials()),
			//grpc2.WithUnaryInterceptor(UnaryInterceptor()),
		),
		grpc.WithTimeout(time.Hour),
	)
	if err != nil {
		t.Error(err)
		return
	}

	client := v1.NewGreeterClient(newClient)

	rsp, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "err"})
	if err != nil {
		fromError := errors.FromError(err)
		t.Error(fromError)
		return
	}

	t.Log(rsp)
}
