package main

import "fmt"

type car struct {
	name   string
	speed  int
	weight float64
}

func main() {
	car1 := car{
		name:   "Ford Mustang 2020",
		speed:  340,
		weight: 2000.2,
	}

	fmt.Println(car1)

	fmt.Printf("\ncar name: %s \n max speed: %d \n weight: %f \n", car1.name, car1.speed, car1.weight)
}
