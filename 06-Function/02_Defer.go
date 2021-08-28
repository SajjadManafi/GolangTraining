package main

import "fmt"

func main() {

	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.

	openFile()
	defer closeFile()
	WriteFile()

	// Opening -> Writing -> Closing

	// Second Example:

	// defer fmt.Println("one")
	// defer fmt.Println("Two")
	// fmt.Println("Three")

	// Three -> Two -> One

	// more -> https://go.dev/blog/defer-panic-and-recover and https://tour.golang.org/flowcontrol/12
}

func openFile() {
	fmt.Println("Opening")
}
func closeFile() {
	fmt.Println("Closing")
}
func WriteFile() {
	fmt.Println("Writing")
}
