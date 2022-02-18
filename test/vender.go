package main

import "fmt"

type Vender int

const  (
	ali Vender = iota
	tengxun
	huawei
)

func main()  {
	fmt.Println(ali)
}
