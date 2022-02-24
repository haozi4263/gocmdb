package main

import (
	"context"
	"fmt"
	"github.com/EDDYCJY/go-grpc-example/pkg/gtls"
	"google.golang.org/grpc"

	"grpc/protos"
	"log"
	"net/http"
	"strings"
)

type EchoService struct {
	protos.UnimplementedEchoServiceServer
}

// unary rpc 一元调用，客户端发送一次请求，服务端一次返回
// server stream rpc, 客户端发送一次请求，服务端多次返回
// client stream rpc，客户度发送多次请求，服务端一次返回
// bidirectional rpc, 双向数据流
func (es *EchoService) UnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	res := "received: " + req.GetReq()
	return &protos.EchoResponse{Rsp: res}, nil
}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("grpc-http: grpc-test"))
	})
	return mux
}

func main() {
	certFile := "/Users/zhanghao/go/src/htgolang-201906/course/20191214/code/gocmdb/grpc/http_grpc/cert/server.pem"
	keyFile := "/Users/zhanghao/go/src/htgolang-201906/course/20191214/code/gocmdb/grpc/http_grpc/cert/server.key"
	tlsServer := gtls.Server{
		CertFile: certFile,
		KeyFile:  keyFile,
	}

	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}

	httpServer := GetHTTPServeMux()

	//创建grpc server
	grpcServer := grpc.NewServer(grpc.Creds(c))
	protos.RegisterEchoServiceServer(grpcServer, &EchoService{}) //绑定

	fmt.Println(5)

	err = http.ListenAndServeTLS(":8080",
		certFile,
		keyFile,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				httpServer.ServeHTTP(w, r)
			}
			return
		}),
	)

}
