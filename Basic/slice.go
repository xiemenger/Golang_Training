package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
	slice1 := array1[2:4:5]             // [start: end : cap] , start是包含的, end是不包含的,
	fmt.Println("slice is ", slice1)    // [3 4]  [start .... end - 1]
	fmt.Println("len is ", len(slice1)) // 2  end - start - 1
	fmt.Println("cap is ", cap(slice1)) // 3  cap - start

	//----- 初始化 ------------------
	// 方法一 自动推导类型
	s1 := []int{1, 2}
	fmt.Println(s1) // [1 2]

	// 方法二 make
	s2 := make([]int, 5, 10) // 长度为5, cap 为10, 如果没有指定cap
	// 那么cap = 长度
	fmt.Println(s2) // [0 0 0 0 0]

	s3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 截取规则 [low: hight : max]  取下标为low, 到high-1的数值, cap = max - low;

	//----- slice与底层数组的关系 ----------------
	s4 := s3[2:5]
	fmt.Println("s4: ", s4) // 2 3 4]
	s4[0] = 666
	fmt.Println("s4: ", s4) // [666 3 4]
	fmt.Println("s3, ", s3) // [0 1 666 3 4 5 6 7 8 9]
	// 底层数组数值也改变了

	//----- 常用方法 ----------------
	// append func, 如果原来的容量不够用, 会自动扩容 (2倍的容量扩容)
	s5 := []int{}
	fmt.Printf("s5: len = %d, cap = %d \n", len(s5), cap(s5)) // s5: len = 0, cap = 0
	s5 = append(s5, 1)
	fmt.Printf("s5: len = %d, cap = %d \n", len(s5), cap(s5)) // s5: len = 1, cap = 1
	s5 = append(s5, 2)
	fmt.Printf("s5: len = %d, cap = %d \n", len(s5), cap(s5)) // s5: len = 2, cap = 2
	s5 = append(s5, 3)
	fmt.Printf("s5: len = %d, cap = %d \n", len(s5), cap(s5)) // s5: len = 3, cap = 4

	// copy func
	srcSlice := []int{1, 2}
	dstSlice := []int{3, 4, 5, 6}

	copy(dstSlice, srcSlice)
	// 会把dstSlice的部分元素覆盖
	fmt.Println("dstSlice: ", dstSlice) // dstSlice:  [1 2 5 6]

	//----- slice as parameter -------
	// 会传递地址, 并非传真实值, 这个和array不一样
	n := 10
	s := make([]int, n)
	fmt.Println("s before InitData: ", s)
	InitData(s)
	fmt.Println("s after InitData: ", s)
	bubbleSort(s)
	fmt.Println("s after bubbleSort: ", s)
}

func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) // 100的随机数
	}
}

func bubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
