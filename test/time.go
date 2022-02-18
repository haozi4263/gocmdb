package main

import (
	"fmt"
	"time"
)

func main()  {
	str := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05 MST"))
	tt, _ := time.Parse("2006-01-02 15:04:05 MST", str)
	fmt.Println(tt)
}
