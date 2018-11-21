package main

import "fmt"

type Human interface { //父interface
	sayhi()
}

type Person interface { //子interface
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
	var p Person

	// 学生struct
	s := &Student{"Diane", 555}
	p = s

	p.sayhi()       // 继承过来的方法  Student [Diane, 555] sayhi
	p.sing("moon ") // Student is singing  moon
}
