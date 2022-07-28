package main

import "fmt"

func main() {
	//① 先声明再赋值
	var x int
	x = 18
	fmt.Println(x)
	var y int
	fmt.Println(y) //声明未赋值，有默认零值
	//② 声明并赋值一行实现
	// var name string = "yan"
	var name = "yan"
	fmt.Println(name)

	//③ 声明并赋值的简洁写法
	name2 := "rain"
	fmt.Println(name2)

	// ④ 声明多个变量
	var (
		a int = 8
		b string
		c bool
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//⑤ 一行声明赋值多个变量
	var name3, age, sex = "yan", 23, false
	fmt.Println(name3, age, sex)
}
