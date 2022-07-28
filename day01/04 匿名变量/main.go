package main

import "fmt"

// 使用场景，调用一些函数，但是又不想返回函数中的某些值
func foo() (int, int) {
	return 100, 2
}
func main() {
	var a, _ = foo()
	fmt.Println(a)

	name := "liuyan"
	firstName := "liu"
	secondtName := "yan"
	fmt.Println(name)
	fmt.Println(firstName)
	fmt.Println(secondtName)

	//_1a := 20
	//fmt.Println(_1a)
}
