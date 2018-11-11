package main

import "fmt"

func main() {
	var ch byte // 字符类型
	var str string

	ch = 'a'

	// String is actually consisted by byte (字符类型)
	// 只是字符与字符之间省略了结束符 \0, 只在string末尾有结束符
	// 所以字符串也是可以一个个读取的str[0]
	str = "abcd"
	fmt.Println(ch)
	fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
}
