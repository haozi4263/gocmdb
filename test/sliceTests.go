package main

import "fmt"

var allns = []map[string]string{}

func main()  {
	allns = make([]map[string]string, 1, 20)
	allns[0]= map[string]string{
		"id":"0",
		"text":"default",
	}
	fmt.Println(allns[0])
}