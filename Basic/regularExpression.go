/**
正则表达式 Go 实现的是RE2 标准
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc  adwer, woeirelk sss"

	// 1) 解释规则
	// 这个函数会解析正则表达式, 如果匹配, 则返回ture, 否则返回false
	exp1 := regexp.MustCompile("a.c")
	if exp1 == nil {
		fmt.Println("regexp error!")
		return
	}

	// 2) 根据规则提取关键信息
	res1 := exp1.FindAllStringSubmatch(str, -1)
	fmt.Println("res1: ", res1) // res1:  [[abc]]

}
