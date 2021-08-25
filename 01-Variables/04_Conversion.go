package main

import "fmt"

type wtf int

var X wtf = 209

func main() {

	// var Q int = X -> we cant do this
	// we must convert wtf to int

	var Q int = int(X)
	fmt.Println(Q)

	//Conversion -> casting in other programming languages

}
