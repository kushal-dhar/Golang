package main

import (
	"fmt"
)

var a int // Declare a variable of type int
type mytype int

var b mytype

func main() {
	a = 10
	b = 20
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	b = mytype(a)
	fmt.Println(b)
}

