package main

import "fmt"

func main() {
	var s string //没有赋值，默认则为空
	fmt.Println(s)
	s = "hello yan!"
	fmt.Println(s)
	//1、索引 字符串[索引]
	fmt.Println(string(s[1]))
	fmt.Println(string(s[6]))

	//2、切片字符串[strat:end]
	fmt.Println(s[0:5])
	fmt.Println(s[5])
	fmt.Println(s[6:11])
	fmt.Println(s[6:])
	fmt.Println(s[:])
	//3、字符串拼接 +
	var s1 = "liu"
	var s2 = "yan"
	fmt.Println(s1 + s2)
	//4、转义符号
	fmt.Println("aaa\ncedv\ncefer")
	//取消特殊作用，比如路径
	var s3 = "D:\\next\\weixin"
	fmt.Println(s3)

	var s4 = "his name is \"rain\""
	fmt.Println(s4)
	//5、多行打印
	/*	fmt.Println("1. 买血")
		fmt.Println("2. 购买武器")
		fmt.Println("3. 生命值")*/

	info := `
	1. 买血
	2. 购买武器
	3. 生命值
	请输入选择：
    `
	fmt.Println(info)

}
