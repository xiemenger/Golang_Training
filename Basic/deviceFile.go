package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Stdout.Close() //会把终端输出关闭, 下面的println就打印不出来了
	fmt.Println("Are you OK?")

	os.Stdout.WriteString("Are you ok??? \n") //往终端屏幕上输出

	// os.Stdin.Close()  关闭后就无法输入
	var a int
	fmt.Println("please input a number")
	fmt.Scan(&a)
	fmt.Println("a = ", a)

	path := "./demo.txt"
	WriteFile(path)
}

func WriteFile(path string) {
	//新建文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完毕, 需要关闭文件
	defer file.close()
	for i := 0; i < 10; i++ {
		// 把"i = 1\n" 存储在buf中
		buf := fmt.Sprintf("i = %d\n", i)
		n, err := file.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
	}
}
