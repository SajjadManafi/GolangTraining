package main

import "fmt"

type car struct {
	name  string
	speed int
}

func main() {

	c1 := car{
		name:  "Bugatti Veyron",
		speed: 450,
	}

	c1.PrintMethod()

	PrintCar(c1)
}

func (c car) PrintMethod() {
	fmt.Println(c.name, c.speed)
}

func PrintCar(c car) {
	fmt.Println(c.name, c.speed)
}
