package main

import (
	"bufio"
	"fmt"
	"io"
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

func ReadFile(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 关闭文件
	defer file.close()

	buf := make([]byte, 1024*2)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF { //文件出错, 同时没有到file end
		fmt.Println("err = ", err)
		return
	}
}

//每次读取一行
func ReadFileLine(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 关闭文件
	defer file.close()

	//新建一个缓冲区, 把内容先放在缓冲区
	r := bufio.NewReader(file)

	for {
		//遇到'\n'结束读取, 但\n 也读取进去了
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = %s", string(buf))
	}

}
