package main

import "fmt"

// 可见性
// 如果像使用别的包的函数, 结构体类型,
// 结构体诚园, 函数名, 类型名, 结构体成员变量名, 首字母必修大写
// 如果首字母是小写, 那只能再自己包里可以用
type Student struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func main() {
	//顺序初始化, 每个成员必须初始化
	var s1 Student = Student{1, "mike", 'm', 18, "Beijing"}
	fmt.Println("s1 = ", s1) // s1 =  {1 mike 109 18 Beijing}

	//部分成员初始化, 自动推导类型的定义
	s2 := Student{name: "Joy", addr: "Nanchang"}
	fmt.Println("s2 = ", s2) // s2 =  {0 Joy 0 0 Nanchang}

	// 指针运算
	var s Student
	var p1 *Student // 指针变量保存s的地址
	p1 = &s

	// 通过指针操作成员 p1.id 和 (*p1).id完全等价
	p1.id = 18
	(*p1).name = "Nacy"
	p1.sex = 'F'
	fmt.Println("s = ", s) // s =  {18 Nacy 70 0 }

	// 通过new申请一个结构体
	p2 := new(Student)
	p2.name = "Diane"
	p2.age = 10
	fmt.Println("p2 = ", *p2) // p2 =  {0 Diane 0 10 }

	// 结构体比较, 只能== or !=
	s3 := Student{1, "mike", 'm', 18, "Beijing"}
	s4 := Student{1, "mike", 'm', 18, "Beijing"}
	s5 := Student{1, "mike", 'm', 18, "Beijing"}
	s6 := Student{1, "Dave", 'm', 18, "Beijing"}
	fmt.Println("s3 == s4", s3 == s4) // true
	fmt.Println("s5 == s6", s5 == s6) // false

	// struct作为参数传递, 会进行值传递, 不是地址传递
	fmt.Println("s1 before changeid:  ", s1) // {1 mike 109 18 Beijing}
	changeid(s1)                             // 改变id
	fmt.Println("s1 after changeid: ", s1)   // {1 mike 109 18 Beijing}
	// 还是一样的

}

func changeid(s Student) {
	// change value
	s.id = 100
}
