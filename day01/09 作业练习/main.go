package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "hello golang"
	//fmt.Println(s[6:])
	//#8.将"北京 上海 广州 深圳"转换为 "北京,上海,广州,深圳"
	s1 := "北京 上海 广州 深圳"
	nameSlice := strings.Split(s1, " ")
	fmt.Println(nameSlice)
	fmt.Println(nameSlice[0])
	var news1 = strings.Join(nameSlice, ",")
	fmt.Println(news1)

	//9.基于字符串操作获取用户名和密码s := "mysql ... -u root -p 123"  ，并判断用户名和密码是否是“yuan”和123，获取布尔值。
	//s3 := "mysql ... -u root -p 123"
	//fmt.Println(s3[13:17])
	//fmt.Println(s3[21:])

	cmd := "mysql ... -u root -p 123"
	cmdSlice := strings.Split(cmd, " ")
	//userName := cmdSlice[3]
	//passWord := cmdSlice[5]
	//fmt.Println(userName, passWord)
	fmt.Println(cmdSlice[3], cmdSlice[5])

	//10.引导用户输入一个名字，判断改名字是否以小写a或者大写A开头，如果是打印true，否则打印false
	var name string = "Aaa"
	//fmt.Scanf("1:%s", &name)
	fmt.Println(strings.HasPrefix(name, "a"))
	fmt.Println(strings.HasPrefix(name, "A"))
}
