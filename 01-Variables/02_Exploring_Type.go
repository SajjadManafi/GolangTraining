package main

import "fmt"

var x string = "How are you?"

func main() {

	fmt.Println(x)
	fmt.Printf("x type : %T \n", x)

	// z = 10
	// we cant do that.

	// go is a STATIC programming language not a DYNAMIC programming language
	// VARIABLE is DECLARED to hold a VALUE of a certain TYPE

	var y string = `see. if we want prtint double quotation ->  ""`

	fmt.Println(y)
	fmt.Printf("a type :%T \n", y)
}
