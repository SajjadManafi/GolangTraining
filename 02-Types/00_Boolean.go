package main

import "fmt"

var b bool
var z bool

func main() {

	// by default bool assign false
	fmt.Println(b)

	b = true
	fmt.Println(b)

	z = 10 < 20
	fmt.Println(z)

	z = 10 == 20
	fmt.Println(z)

	a := 42
	b := 90

	z = a != b
	fmt.Println(z)

}
