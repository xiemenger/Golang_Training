package main

type Human interface {
	SayHi()
}

type Student struct {
	name string
	id int
}

func (s *Student) SayHi(){
	fmt.Printf("Student [%s, %d]  sayhi \n", s.name, s.id)
}

type Teacher sturct{
	addr string
	group string
}

func (t *Teacher) SayHi(){
	fmt.Printf("Teacher [%s, %s]  sayhi \n", t.addr, t.group)
}

type MyStr string

func (m *MyStr) SayHi(){
	fmt.Printf("Mystr [%s] sayhi \n", *m)
}

// 定义一个普通函数, 函数的参数为接口类型
func WhoSayHi(i Human){
	i.SayHi()
}

func main(){
	// 定义接口类型的变量
	var h Human 

	// 只要实现此接口方法的类型, 那么这个类型的变量就可以给h赋值
	// Student, Teacher, Mystr都实现了Sayhi, 那都可以付给i. 都是human 类型
	s := &Student("mike", 666)
	h = s
	h.SayHi()

	t := &Teacher("bj", "Go")
	h = t
	h.SayHi()

	var m MyStr := "Hello Yo"
	h = m
	h.SayHi()

	// 调用同一函数, 不同的表现, 多态, 多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&m)

	// 定义一个切片, 
	x := make ([]Human, 3)
	x[0] = s
	x[1] = t
	x[2] = &m

	// range 返回两个值, 第一个是下标i, 第二个是value
	for _, v := range (x){
		v.sayhi()
	}
}


 