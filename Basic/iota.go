package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a = %d, b = %d , c = %d \n", a, b, c)

	const d = iota
	fmt.Printf("d = %d\n", d)

	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d , c1 = %d ", a1, b1, c1)

	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i = %d,\n j1, j2, j3 = %d, %d, %d\n k = %d \n", i, j1, j2, j3, k)
}
