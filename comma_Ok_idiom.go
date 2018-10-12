package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 10
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}

/**
10 true
0 false
*/
