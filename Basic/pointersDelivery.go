package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Printf("a = %d,  b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %d,  b = %d\n", a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
