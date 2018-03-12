package main

import (
	"fmt"
)

func main() {
	total := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		//fmt.Println("value: ", v, sum)
		sum += v
	}
	return sum
}

