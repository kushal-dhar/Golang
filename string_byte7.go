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
}

