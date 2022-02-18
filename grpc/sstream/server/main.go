package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"net"
)

type EchoService struct {
	protos.UnimplementedEchoServiceServer
}

// server stream rpc, 客户端发送一次请求，服务端多次返回
func (es *EchoService)SStreamEcho(req *protos.EchoRequest, stream protos.EchoService_SStreamEchoServer) error {
	fmt.Println("received: ", req.GetReq())
	for i := 1; i <= 3; i ++ {
		res := fmt.Sprintf("%d send: ",i) + req.Req
		stream.Send(&protos.EchoResponse{Rsp: res})
	}
	return nil
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
