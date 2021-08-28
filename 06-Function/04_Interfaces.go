package main

import (
	"fmt"
	"math"
)

// Polymorphism is the ability of a message to be displayed in more than one form.
// Polymorphism is considered as one of the important features of Object-Oriented
// Programming and can be achieved during either at runtime or compile time.
// Golang is a light-Object Oriented language and supports polymorphism through interfaces only.

// for more :
//			https://www.geeksforgeeks.org/polymorphism-in-golang/
//			https://gobyexample.com/interfaces
//			https://tour.golang.org/methods/9

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return (r.height + r.width) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func measure(g geometry) {
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func PrintType(g geometry) {
	switch g.(type) {
	case circle:
		fmt.Println("circle, radius:", g.(circle).radius)
	case rect:
		fmt.Println("rectangle, width:", g.(rect).width, " height:", g.(rect).height)
	}
}

func main() {

	r := rect{width: 6, height: 9}
	c := circle{radius: 3}

	PrintType(r)
	measure(r)

	fmt.Println("")

	PrintType(c)
	measure(c)

}
