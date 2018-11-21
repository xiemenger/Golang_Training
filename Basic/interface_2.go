package main

import "fmt"

type Human interface { //子集
	sayhi()
}

type Person interface { //超集
	Human // 匿名字段, 继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (s Student) sayhi() {
	fmt.Printf("Student [%s, %d] sayhi\n", s.name, s.id)
}

func (s Student) sing(lrc string) {
	fmt.Println("Student is singing ", lrc)
}

func main() {
	//定义接口类型
	var p Person // 超集

	// 学生struct
	s := &Student{"Diane", 555}
	p = s

	p.sayhi()       // 继承过来的方法  Student [Diane, 555] sayhi
	p.sing("moon ") // Student is singing  moon

	// ---------接口转换----------
	// 超集可以转换为子集, 但是反过来不可以.
	// var child Person // 超集
	// var father Human // 被继承者集

	// p = h  // err
	//father = child
	//father.sayhi()

	// --------- 空接口 ----------
	// 空接口就是没有方法的接口, 是个万能类型, 可以保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i) // i =  1

	i = "abc"
	fmt.Println("i = ", i) // i =  abc

	slice := make([]interface{}, 3)
	slice[0] = 1
	slice[1] = "hello"
	slice[2] = Student{"Qiqi", 19}

	for i, v := range slice {
		if _, ok := v.(int); ok == true {
			fmt.Printf("x[%d] is int\n", i)
		} else if _, ok := v.(string); ok == true {
			fmt.Printf("x[%d] is string\n", i)
		} else if _, ok := v.(Student); ok == true {
			fmt.Printf("x[%d] is student\n", i)
		}
	}
	// x[0] is int
	// x[1] is string
	// x[2] is student

	//用switch
	for i, v := range slice {
		switch typeValue := v.(type) {
		case int:
			fmt.Printf("x[%d] is int\n", i)
		case string:
			fmt.Printf("x[%d] is string\n", i)
		case Student:
			fmt.Printf("x[%d] is student, studnet name is %s\n", i, typeValue.name)
		}
	}

}
