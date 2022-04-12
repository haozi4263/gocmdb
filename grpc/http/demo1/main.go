package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	// 用户提交的数据http内容-->go代码转换成http.Request
	w.Write([]byte("home"))
}

type Help struct{}
func (h *Help)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "help")
}


func main()  {
	addr := ":8888"
	http.HandleFunc("/home",Home)
	http.Handle("/help", new(Help))

	http.ListenAndServe(addr, nil)
}