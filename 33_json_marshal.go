package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		"Kushal",
		"Dhar",
		25,
	}

	p2 := Person{
		"Kunal",
		"Dhar",
		31,
	}
	people := []Person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

