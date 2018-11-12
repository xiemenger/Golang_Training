package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 设置种子, 只需要一次
	// 如果种子参数一样, 每次运行程序产生的随时间数都一样
	rand.Seed(time.Now().UnixNano()) // 以当前系统时间当为seed

	for i := 0; i < 5; i++ {
		fmt.Println("rand is ", rand.Int())
	}

	fmt.Println("rand is ", rand.Intn(100)) // 限定100以为的数

}
