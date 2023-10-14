package service

import (
	"context"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/internal/pkg/errs"
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

	//return &v1.HelloReply{Message: "Hello " + g.Hello,}, nil
	return nil, errs.New(2000, "error123")
}
