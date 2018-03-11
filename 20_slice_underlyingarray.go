package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	y := append(x, 11, 12, 13, 14, 15) // a new underlying array is allocated
	fmt.Println(y)
	fmt.Println(x)

	z := append(x[:2], x[4:]...)
	fmt.Println(z)
	fmt.Println(x)
}

