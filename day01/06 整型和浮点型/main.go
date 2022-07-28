package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1、整型
	//int8 :一个字节 范围2的8次方 [-128-127]
	//uint8:一个字节 范围2的8次方 [0-255]
	//uint16:两个字节 范围2的8次方 [0-65535]
	var x uint16 = 65535
	fmt.Println(x)
	var y int16 = -32768 //-32768到32767
	fmt.Println(y)

	//2、浮点型
	var f1 float32 = 3.1415
	var f2 float64 = 3.1415
	fmt.Println(f1, reflect.TypeOf(f1))
	fmt.Println(f1, reflect.TypeOf(f2))

	var f3 = 2e10 //2的10次方
	fmt.Println(f3, reflect.TypeOf(f3))

	// 使用场景，登录账户是否正确
	name := "yan"
	c2 := name == "root"
	fmt.Println(c2, reflect.TypeOf(c2))
}
