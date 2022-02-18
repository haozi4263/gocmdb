package main

import "fmt"

func test() func()  {
	x := 100
	fmt.Println("run text global")
	return func() {
		fmt.Println("run func.. x=",x)
	}
}

// 外部引用函数参数局部变量
func add(sum int) func(int) int {
	return func(i int) int {
		sum += i
		return sum
	}
}

func main()  {
	fn := test()
	fn()

	sum := add(10)
	i := sum(10)
	fmt.Println(i)
}