package main

import (
	"fmt"
)

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0])

	n := bs[0]
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}

