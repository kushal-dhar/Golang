package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	b := []byte(s)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\t", s[i])
	}
	fmt.Println()

	for i, v := range s {
		fmt.Println(i, v)
	}
}

