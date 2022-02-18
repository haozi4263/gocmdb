package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
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
	fmt.Println(1, res)
	return &protos.EchoResponse{Rsp: res}, nil
}


func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http: go-http"))
	})
	return mux
}

func main() {
	server := grpc.NewServer()                               //创建grpc server
	protos.RegisterEchoServiceServer(server, &EchoService{}) //绑定
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ProtoMajor", r.ProtoMajor)
		if r.ProtoMajor == 2 && strings.Contains(r.Proto, "HTTP/2.0") {
			server.ServeHTTP(w, r)
		} else {
			mux := GetHTTPServeMux()
			mux.ServeHTTP(w, r)
		}
		return
	}),
	)

}

