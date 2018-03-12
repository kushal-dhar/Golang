package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr := truck{
		vehicle{
			6,
			"white",
		},
		false,
	}

	car := sedan{
		vehicle{
			4,
			"red",
		},
		true,
	}
	fmt.Println(tr)
	fmt.Println(car)
	fmt.Printf("%T\t%T\n", tr.vehicle, tr)
	fmt.Printf("%T\t%T", car.vehicle, car)
}

