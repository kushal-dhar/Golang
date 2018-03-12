package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("Anonymous func ran with value ", x)
	}(42)
}

