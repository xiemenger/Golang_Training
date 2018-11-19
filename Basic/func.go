package main

import "fmt"

// main func (receiver ReceiverType) funcName (parameterName type) (returnType)

// add function
// 面向过程
func Add01(a, b int) int {
	return a + b
}

// 面向对象, 方法: 给某个类型绑定一个函数
type long int

// tmp叫接受者, 接受者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}

type Person struct {
	name   string
	gender byte
	age    int
}

type Student struct {
	Person
	id   int
	addr string
}

//接受者为普通变量, 非指针, 值语义, 一份拷贝
func (tmp Person) PrintInfo() {
	fmt.Println("tmp : ", tmp)
}

// 接受者为指针变量, 引用传递
func (p *Person) SetInfo(n string, g byte, a int) {
	p.name = n
	p.gender = g
	p.age = a
}

func (p *Person) PrintAddr() {
	fmt.Println("p : ", p)
}

func main() {
	var res int
	res = Add01(1, 2)           //普通函数调用方法
	fmt.Println("res  = ", res) // res  =  3

	// 定义一个变量
	var a long = 2
	var res2 long

	// 调用方法的格式: 变量名.函数(所需参数)
	res2 = a.Add02(3)
	fmt.Println("res2 = ", res2) // res2 =  5

	p := Person{"Mike", 'M', 20}
	p.PrintInfo() // tmp :  {Mike 77 20}

	var p2 Person
	(&p2).SetInfo("Yoyo", 'F', 22)
	p2.PrintInfo() // tmp :  {Yoyo 70 22}

	//---------继承---------------------
	// 学生继承Person字段, 成员和方法都会继承
	s := Student{Person{"Dave", 'f', 20}, 666, "qingdao"}
	s.PrintAddr() // tmp :  &{Dave 102 20}

	// --------重写---------------------
	// Student 也实现了一个方法, 这个方法和Person方法同名, 这种方法叫做重写
	// 就近原则, 先找本作用域的方法, 找不到的话在找继承的方法
	s.PrintAddr() // tmp :  &{Dave 102 20}

	// --------方法值 ---------------
	pFunc := s.PrintAddr // 这个就是方法值, 调用函数时候, 无需再传递接受者, 隐藏了接受者
	pFunc()              // 等价于s.PrintAddr()

	// -------- 方法表达式 --------------
	f := (*Person).PrintAddr
	f(&p) // 显示把接受者传递过去

	f2 := (Person).PrintInfo
	f2(p2)

}

// student重写
func (s *Student) PrintAddr() {
	fmt.Println("s : ", s)
}
