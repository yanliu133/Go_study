package main

import (
	"fmt"
	"strings"
)

func main() {
	//大小写转换
	var name = "Yuan"
	var newName = strings.ToUpper(name)
	fmt.Println(newName)
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))

	//判断以什么开头是否正确
	var s = "rain yuan alvin"
	fmt.Println(strings.HasPrefix(s, "rain"))
	//判断以什么结尾是否正确
	fmt.Println(strings.HasSuffix(s, "in"))
	//判断以什么中间是否正确
	fmt.Println(strings.Contains(s, "yuan"))

	//去两端空格/换行符
	username := "   yuan   "
	fmt.Println(strings.Trim(username, " "))

	//去掉左右空格，判断是否是正确的字符
	username = strings.TrimSpace(username)
	username = strings.TrimLeft(username, "")
	fmt.Println(username == "yuan")

	//index :索引
	var s2 = "rain yuan alvin"
	fmt.Println(strings.Index(s2, "alex")) //没有，返回-1
	var cmd = "mysql ... -u root -p 123"   //场景，数据库确认账户和密码等
	fmt.Println(cmd)

	// 分割 拼接
	var s3 = "rain yuan alvin"
	nameSlice := strings.Split(s3, " ")
	fmt.Println(nameSlice)
	fmt.Println(nameSlice[0])
	fmt.Println(nameSlice[1])
	fmt.Println(nameSlice[2])

	var newStr = strings.Join(nameSlice, ",")
	fmt.Println(newStr)
}
