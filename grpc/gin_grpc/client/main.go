package main

import (
	"github.com/golang/protobuf/proto"

	pb "grpc/gin_grpc/proto"
	"io/ioutil"
	"net/http"
	"fmt"
)

func main()  {
	resp, err := http.Get("http://localhost:8080/proto")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			teacher := &pb.Teacher{}
			proto.UnmarshalMerge(body, teacher)
			fmt.Println(teacher.Name, teacher.Course)
		}
	}
}
