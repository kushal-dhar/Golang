package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[:])

	x = append(x, 10, 12, 13, 14, 15)
	fmt.Println(x)

	y := []int{100, 102, 103, 104}
	x = append(x, y...)
	fmt.Println(x)
}

