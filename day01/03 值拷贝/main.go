package main

import "fmt"

func main() {
	// 1、值拷贝
	var x = 10
	var y = x //x,y 对应的是完全不同的地址空间
	x = 20
	fmt.Println(x)
	fmt.Println(y)

	//2、
	var a = 1 + 1
	fmt.Println(a)
	var b = x * y
	fmt.Println(b)
	var c = x + y
	fmt.Println(c)

	//
	fmt.Println("yan")
	name := "yan"
	fmt.Println(name + ".liu")

	var (
		p int = 10
		q int = 20
	)
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)
	fmt.Println(p % q)

}
