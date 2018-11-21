package main

import (
	"errors"
	"fmt"
)

func main() {
	// ----- 两种方法来产生Error 类型-------------
	// 方法一, 用ftm.Errorf
	//var err error := fmt.Errorf("%s", "This is a normal error")
	err := fmt.Errorf("%s", "This is a normal error")
	fmt.Println("err is : ", err) // err is :  This is a normal error

	//方法二, 用errors.New(string)
	err2 := errors.New("This is another error ")
	fmt.Println("err2 is : ", err2) // err2 is :  This is another error

	// ----- error的应用 ----------------------
	result, err := myDiv(10, 2)
	if err != nil {
		fmt.Println("nice catch! ", err)
	} else {
		fmt.Println("The result is ", result)
	}
}

func myDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("The divident cannot be zero")
	} else {
		result = a / b
	}
	return result, err
}
