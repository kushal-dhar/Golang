package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type sa struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		name: "Kushal",
		age:  25,
	}

	sa1 := sa{
		person: person{
			name: "James Bond",
			age:  35,
		},
		ltk: true,
	}

	fmt.Println(p1.name, p1.age)
	fmt.Println(sa1.name, sa1.age, sa1.ltk)
	fmt.Println(sa1)
}

