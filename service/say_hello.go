package service

import (
	pb "github.com/test/proto/example"

	"golang.org/x/net/context"
)

// SayHello implements helloworld.Greeter SayHello rpc service.
func (GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	if req.Name == "" {
		return nil, NewError(MissingNameParam)
	}

	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}
