package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	sides float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.sides * s.sides
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{
		4,
	}

	sq := square{
		4,
	}

	info(c1)
	info(sq)
}

