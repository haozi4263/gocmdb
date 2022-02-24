package main

import (
	"context"
	"fmt"
	"github.com/EDDYCJY/go-grpc-example/pkg/gtls"
	"google.golang.org/grpc"
	"grpc/protos"
	"log"
)



func main()  {
	//链接grpc服务
	tlsClient := gtls.Client{
		ServerName: "grpc-test", // 需要和证书里面的Common Name保持一致
		CertFile: "/Users/zhanghao/go/src/htgolang-201906/course/20191214/code/gocmdb/grpc/http_grpc/cert/server.pem",
	}

	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(c)) // 使用安全的链接
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := protos.NewEchoServiceClient(conn)

	res ,err := client.UnaryEcho(context.Background(), &protos.EchoRequest{
		Req: "grpc-rpc",
	})
	fmt.Println("resp: ", res.GetRsp())

}
