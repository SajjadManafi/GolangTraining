package main

import "fmt"

func main() {

	x := 12
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := 90.23245
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// this is doesnt work -> var z int8 = 128

	//rule of thumb: just use int
	//rule of thumb: just use float64

	// uint8 		-> 	the set of all unsigned  8-bit integers (0 to 255)
	// int8  		-> 	the set of all signed  8-bit integers (-128 to 127)
	//float64		->	the set of all IEEE-754 64-bit floating-point numbers
	//complex128	-> 	the set of all complex numbers with float64 real and imaginary parts
	//byte        	->	alias for uint8
	//rune        	->	alias for int32

	// for more -> https://golang.org/ref/spec#Numeric_types
}
