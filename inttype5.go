package main

import (
	"fmt"
)

var x int
var y float64

func main() {
	x = 2
	y = 10.012
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%d\n", x)
	fmt.Printf("%T\n", y)
}

