package main

import (
	"Golang_Training/Advanced"
	"fmt"
)

/*
每个源文件都可以包含一个init函数, 这个init函数自动被go运行框架调用, 一般是变量的初始化, 配置之类的.
,在main执行之前就会调用.
*/

func main() {
	fmt.Printf("var1 : %d, var2 : %d \n", Advanced.Var1, Advanced.Var2)
	// var1 : 100, var2 : 200
}
