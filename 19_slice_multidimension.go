package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Juice"}
	fmt.Println(jb)
	mp := []string{"miss", "moneypenny", "strawberry", "vodka"}
	fmt.Println(mp)

	md := [][]string{jb, mp}
	fmt.Println(md)
}

