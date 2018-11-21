/**
运行时panic异常一旦被引发就会导致程序崩溃, 这当然不是我们所想要的, 因为
谁也不能保证程序不会发生任何运行时错误

不过, Go语言为我们提供了装用于"拦截"运行时的panic的内建函数 recover
它可以使当前的程序从运行时panic的状态中恢复并重新获得流程控制权
*/

package main

import "fmt"

func main() {
	// 设置recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	try1()
	try4(20) // panic: runtime error: index out of range
	try3()   //不会运行, 因为try4有panic, 程序会终端,

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
