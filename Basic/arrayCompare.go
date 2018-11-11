package main

import "fmt"

func main() {

	// 支付比较, 只支持 == 或 != , 比较是不是每一个元素都一样
	// 2个数组比较, 类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}

	fmt.Println("a == b ? ", a == b) // true
	fmt.Println("a == c ? ", a == c) // false
	var d [5]int
	d = a
	fmt.Println("d = ", d) // d =  [1 2 3 4 5]
}
