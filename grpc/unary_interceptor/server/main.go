package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"context"
	"net"
	"time"
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


func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	start := time.Now()
	res, err := handler(ctx, req)
	end := time.Now()
	fmt.Printf("UnaryInterceptor rpc-name: %s, start: %s end: %s, req:%v rsp:%v\n",
		info.FullMethod , start.Format(time.RFC822), end.Format(time.RFC822),  req, res)
	return res, err
}


func UnaryInterceptorEx(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	start := time.Now()
	res, err := handler(ctx, req)
	end := time.Now()
	fmt.Printf("UnaryInterceptorEx rpc-name: %s, start: %s end: %s, req:%v rsp:%v\n",
		info.FullMethod , start.Format(time.RFC822), end.Format(time.RFC822),  req, res)
	return res, err
}


func main()  {
	//rpcs := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor)) //创建grpc server
	rpcs := grpc.NewServer(grpc.ChainUnaryInterceptor(UnaryInterceptor, UnaryInterceptorEx)) //创建grpc server
	protos.RegisterEchoServiceServer(rpcs, &EchoService{}) //绑定
	lis, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
