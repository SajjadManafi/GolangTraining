package main

import "fmt"

func main() {

	x := Returner(10)
	fmt.Println(x)
	fmt.Printf("%T\n", x) // func() int

	y := x()
	fmt.Println(y)        // 1000
	fmt.Printf("%T\n", y) // int

	// func() int --x()--> int

	fmt.Println(Returner(23)()) // 2300
}

func Returner(number int) func() int {
	return func() int {
		return number * 100
	}
}
