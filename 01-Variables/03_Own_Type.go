package main

import "fmt"

// create your own type
type wtf int

var a int = 40

// declare variable and assign a value
var b wtf = 90

func main() {

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
