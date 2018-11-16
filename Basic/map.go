package main

import "fmt"

func main() {
	// ----- make Map --------
	m1 := make(map[int]string) // key 是int, value是string
	// key特点是: 不会有重复
	// 不能用slice, array作为key.
	// value可以是任何类型
	fmt.Print("m1: ", m1, "len : ", len(m1), "\n") // m1: map[]len : 0

	//还可以指定容量
	m2 := make(map[int]string, 10)
	fmt.Print("m2: ", m2, "len : ", len(m2), "\n") // m2: map[]len : 0
	m2[1] = "mike"
	m2[2] = "go"
	fmt.Print("m2: ", m2, "\nlen : ", len(m2), "\n")
	// [1:mike 2:go]
	// len : 2

	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	// 判断一个值是否存在
	// 第一个返回值为key所对应的value, 第二个返回值为key是否存在
	value, ok := m[1]
	if ok {
		fmt.Println("m[1] = ", value) // m[1] =  mike
	} else {
		fmt.Println("key is not existed")
	}

	// 删除key
	delete(m, 1)
	fmt.Println("m = ", m) // m =  map[2:yoyo 3:go]

	test(m)
	fmt.Println("m = ", m) // m =  map[3:go]
}

func test(m map[int]string) {
	delete(m, 2)
}
