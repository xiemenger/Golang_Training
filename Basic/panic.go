/**
panic 与error的不同在于,
panic是不可恢复的错误状态, 比如数组访问越界, 空指针应用, 栈溢出等等
*/
package main

import "fmt"

func main() {
	try1()
	//try2()
	try3() // 不会运行, 因为try2已经panic, 程序终端

	try4(20) // panic: runtime error: index out of range
}

func try1() {
	fmt.Println("Test1111111111111111")
}

func try2() {
	fmt.Println("Test2222222222222222")
	//调用panic 函数, 导致整个程序中断
	panic("This is a panic test")
}

func try3() {
	fmt.Println("Test3333333333333333")
}

func try4(x int) {
	var a [10]int
	a[x] = 111 //当x为20的时候, 导致数组越界, 产生一个panic, 程序崩溃

}
