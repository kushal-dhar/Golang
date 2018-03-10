package main

import (
	"fmt"
)

var a int   // Declare a variable of type int
var b string  // Declare a variable of type string
var c int  // Declare a variable of type int

func main() {
	fmt.Print(a)
	b := "Hello world"
	fmt.Printf("\n%s\n", b)
	c = 12
	fmt.Println(c)
}

