package main

import "fmt"

// untyped const
const hey = "hey!"

// typed const
const hello string = "Hello..."

const x = 13

const (
	y = 45.34
	a = "z"
)

func main() {

	//consts:
	//a simple, unchanging value
	//Only exist at compile time.

	fmt.Println(hey)
	fmt.Printf("%T\n", hey)

	fmt.Println(hello)
	fmt.Printf("%T\n", hello)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y, a)

}
