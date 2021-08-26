package main

import "fmt"

func main() {

	fmt.Println(true && true)   // true
	fmt.Println(false && true)  // flase
	fmt.Println(false && false) // flase
	fmt.Println(true || true)   // true
	fmt.Println(false || true)  // true
	fmt.Println(false || false) // flase
	fmt.Println(!true)          // flase
	fmt.Println(!false)         // true

	if message := "Hey"; 10 < 20 && message == "Hi" {
		fmt.Println("not Print!")
	} else if message == "Hey" {
		fmt.Printf("Hey")
	}

	// for more -> https://golang.org/ref/spec#Logical_operators

}
