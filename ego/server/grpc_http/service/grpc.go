package service

import (
	"context"

	"ego/helloworld"

	"github.com/gotomicro/ego/server/egrpc"
)

func GrpcServer() *egrpc.Component {
	hello := egrpc.Load("server.grpc").Build()
	helloworld.RegisterGreeterServer(hello.Server, &Greeter{server: hello})
	return hello
}

type Greeter struct {
	server *egrpc.Component
	helloworld.UnimplementedGreeterServer
}

// SayHello ...
func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{
		Message: "Hello grpc_http for grpc" + g.server.Address() + request.GetName(),
	}, nil
}
