package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"io"
	"net"
	"strings"
)

type EchoService struct {
	protos.UnimplementedEchoServiceServer
}

// server stream rpc, 客户端发送多次请求，服务端一次返回
func (es *EchoService)CStreamEcho(stream protos.EchoService_CStreamEchoServer) error {
	var i int
	strs := []string{}
	for{
		req, err := stream.Recv() //不断接受客户端发送请求
		if err == io.EOF{ //接受完成
			res := strings.Join(strs, "\n")
			return stream.SendAndClose(&protos.EchoResponse{Rsp: res}) //汇总返回给客户端
		}
		if err != nil {
			continue
		}
		i++
		strs = append(strs,fmt.Sprintf("%d count recv: ", i) + req.GetReq())
		fmt.Println(req.GetReq())

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
