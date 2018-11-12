package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	modify(&a) // [100 2 3 4 5] 有变化
	fmt.Println(a)
}

/*
	注意!!!
	数组作为函数参数, 它是值传递
	实参数组的每个元素给型参数组拷贝一份
	但是用指针/地址, 就会有改变
*/
func modify(a *[5]int) {
	a[0] = 100
}
