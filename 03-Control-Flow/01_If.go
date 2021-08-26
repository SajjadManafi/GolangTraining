package main

import "fmt"

func main() {

	x := 4
	if x == 4 {
		fmt.Println("The number is 4")
	}

	if !true {
		fmt.Println(false)
	}

	if 2 <= 5 {
		fmt.Println("2 <= 5")
	}

	if y := 13; y%2 == 0 {
		fmt.Println("y is even")
	} else {
		fmt.Println("y is odd")
	}

	b := "Hmmm"
	if b == "Hello" {
		fmt.Println("Hello. How are you?")
	} else if b == "Hey" {
		fmt.Println("Hey wassap")
	} else {
		fmt.Println("What?")
	}

	// for more -> https://golang.org/ref/spec#If_statements
}
