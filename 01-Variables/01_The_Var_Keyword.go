package main

import (
	"fmt"

	"github.com/SajjadManafi/GolangTraining/test"
)

// global in this file
var globalY = 11

//global for package (we can use this var in filesthats import this package)
var GlobalY = 12

func main() {
	var x = 42
	fmt.Println(x)
	fmt.Println(globalY)

	// use variable in ../test/variables.go
	fmt.Println(test.GlobalVar)

	//var can be used global and local declaration.
}
