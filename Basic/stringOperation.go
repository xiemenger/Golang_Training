package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Contains
	fmt.Println(strings.Contains("Hellogo", "llo")) // true

	// Joins
	s := []string{"hello", "boy", "going"}
	fmt.Println(strings.Join(s, " ")) // hello boy going

	// Index
	fmt.Println(strings.Index("abcdedf", "cd")) // 2
	fmt.Println(strings.Index("abcdefg", "gf")) // -1

	// Repeat
	fmt.Println(strings.Repeat("abc", 3)) // abcabcabc

	// Split
	buf := "hello@abcd@go"
	slice := strings.Split(buf, "@")
	for _, v := range slice {
		fmt.Println(v)
	}

	// 转换为字符串后追加到字节数组
	// Append
	c := make([]byte, 0, 1024)
	c = strconv.AppendBool(c, true)

	// 第二个数为要追加的数, 第3个为指定10进制方式追加
	c = strconv.AppendInt(c, 1234, 10)

	// 转化为string之后再打印
	fmt.Println("c: ", string(c)) //c:  true1234

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false) // 把bool false转为string类型 "false"
	fmt.Println("str : ", str)      // str :  false

	// 整形转字符串, 常用
	str = strconv.Itoa(6666)
	fmt.Printf("%T\n", str)

}
