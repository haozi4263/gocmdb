package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"

	"ego/server/grpc_http/service"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().Serve(
		service.GrpcServer(),
		service.HttpServer(),
	).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

// grpcurl -d '{"name":"jude"}' -plaintext 127.0.0.1:9005 helloworld.Greeter.SayHello
