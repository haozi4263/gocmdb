package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"context"
	"net"
)

type EchoService struct {
	protos.UnimplementedEchoServiceServer
}

// unary rpc 一元调用，客户端发送一次请求，服务端一次返回
// server stream rpc, 客户端发送一次请求，服务端多次返回
// client stream rpc，客户度发送多次请求，服务端一次返回
// bidirectional rpc, 双向数据流
func (es *EchoService)UnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	res := "received: " + req.GetReq()
	fmt.Println(res)
	return  &protos.EchoResponse{Rsp:res}, nil
}

func main()  {
	rpcs := grpc.NewServer() //创建grpc server
	protos.RegisterEchoServiceServer(rpcs, &EchoService{}) //绑定
	lis, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
