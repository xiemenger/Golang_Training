package main

import "fmt"

func main() {
	var a int
	fmt.Println("please input for a")

	// program will wait for the user to input
	fmt.Scan(&a)

	fmt.Println("please input for a gain")
	fmt.Scanf("%d", &a)
	fmt.Println("a = ", a)
}
