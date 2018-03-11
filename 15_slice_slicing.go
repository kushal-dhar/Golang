package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1:5])
	fmt.Println(x[:])

	for i := 0; i < 5; i++ {
		fmt.Println(x[i])
	}
}

