package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains
	fmt.Println(strings.Contains("Hellogo", "llo")) // true

	// Joins
	s := []string{"hello", "boy", "going"}
	fmt.Println(strings.Join(s, " ")) // hello boy going

	// Index
	fmt.Println(strings.Index("abcdedf", "cd")) // 2
	fmt.Println(strings.Index("abcdefg", "gf")) // -1

	// Repeat
	fmt.Println(strings.Repeat("abc", 3)) // abcabcabc

	// Split
	buf := "hello@abcd@go"
	slice := strings.Split(buf, "@")
	for _, v := range slice {
		fmt.Println(v)
	}
}
