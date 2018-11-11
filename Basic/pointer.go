package main

import "fmt"

func main() {

	// 保存某个变量的地址, 需要用类型指针  *int  保存 int 的地址
	var a int = 10
	var p *int
	p = &a // 方法一
	fmt.Println("a is : \n", a)
	fmt.Printf("p is type:  %T\n", p)

	*p = 200
	fmt.Println("a is updated to : ", a)

	p2 := new(int) // 方法二 但需要重要的是 需要指向什么类型的指针, 比如这里是int
	p2 = &a
	*p2 = 666
	fmt.Println("a is updated to : ", a)
}
