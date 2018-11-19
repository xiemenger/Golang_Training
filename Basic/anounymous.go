package main

import "fmt"

type Person struct {
	name   string
	gender byte
	age    int
}

type Student struct {
	Person // 匿名字段:  只有类型, 没有名字. 即继承了Person里的成员
	// 默认Student就包含一个Person内所有的elements.
	// 实现代码复用
	id   int
	addr string
}

type Person1 struct {
	name   string
	gender byte
	age    int
}

type Student1 struct {
	Person1 // 匿名字段:  只有类型, 没有名字. 即继承了Person里的成员
	// 默认Student就包含一个Person内所有的elements.
	// 实现代码复用
	id   int
	addr string
	name string
}

func main() {
	// 初始化1
	var s1 Student = Student{
		Person{"Dave", 'M', 20},
		1,
		"Shanghai",
	}
	fmt.Println("s1 = ", s1) // s1 =  {{Dave 77 20} 1 Shanghai}

	// 初始化2  自动推导类型
	s2 := Student{
		Person{"JAY", 'M', 20},
		1,
		"Beijing",
	}
	fmt.Printf("s2 = %+v\n", s2) // %+v 打印key, value pair
	// s2 = {Person:{name:JAY gender:77 age:20} id:1 addr:Beijing}$

	// same elements
	// 见Person1, Student1都有name这个element.
	var s3 Student1
	s3.name = "Sophie" // 默认规则, 如果能再本操作与找到此element, 就操作此成员.
	// 如果没有找到, 就照继承的字段.
	s3.gender = 'F'
	s3.age = 19
	fmt.Printf("s3 = %+v\n", s3) // s3 = {Person1:{name: gender:70 age:19} id:0 addr: name:Sophie}

	var s4 Student1
	s4.Person1.name = "Yoyo"
	fmt.Printf("s4 = %+v\n", s4) // s4 = {Person1:{name:Yoyo gender:0 age:0} id:0 addr: name:}
}
