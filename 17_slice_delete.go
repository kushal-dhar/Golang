package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[:])

	y := []int{100, 102, 103, 104}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:3], x[4:]...)
	fmt.Println(x)
}

