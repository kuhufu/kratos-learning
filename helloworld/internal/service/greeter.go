package service

import (
	"context"
	"errors"
	"helloworld/api/errs"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	_, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}

	if in.Name == "err" {
		return nil, errs.ErrorBusiness("业务错误").
			WithCause(errors.New("cause error")).
			WithMetadata(map[string]string{
				"user": "kuhufu",
				"age":  "11",
			})
	}

	return &v1.HelloReply{Message: "Hello " + in.Name}, nil

}
