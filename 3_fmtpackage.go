package main

import (
	"fmt"
)

var a int // Declare a variable of type int
var b string = "Hello there"
var c bool // Declare a variable of type bool

func main() {
	f := "hello Kushal"
	g := `Looks like you are having fun with golang. how's the language so far? :)`
	fmt.Println(a)
	fmt.Printf("%v\n", b)
	fmt.Println(c)

	s := fmt.Sprint(f, g)
	fmt.Printf("%v\n", s)
	s2 := fmt.Sprintf("%s\t%s\n", f, g)
	fmt.Printf("%s", s2)
}

