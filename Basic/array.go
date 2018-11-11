package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := [5]int{6, 7}
	fmt.Println(b) // [1 2 3 4 5]

	c := [5]int{2: 20, 4: 40}
	fmt.Println(c) // [0 0 20 0 40]
}
