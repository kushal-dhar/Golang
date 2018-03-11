package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := struct {
		name string
		age  int
	}{
		name: "Kushal",
		age:  25,
	}

	fmt.Println(p1.name, p1.age)
}

