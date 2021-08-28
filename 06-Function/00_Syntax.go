package main

import "fmt"

func main() {

	Hello()

	sayHello("Bob")

	text := Hey("Sajjad", "Iran")
	fmt.Println(text)

	x, y := BiggerThan(92, 100)
	fmt.Println(x, y)

}

// func (receiver) identifier(parameters) (returns) {code}
func Hello() {
	fmt.Println("Hello")
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func Hey(name, country string) string {
	s := fmt.Sprintf("Hello. my name is %s and i'm from %s", name, country)
	return s
}

func BiggerThan(a int, b int) (int, bool) {
	if a >= b {
		return a, true
	} else {
		return b, false
	}
}

// Go's return values may be named. If so,
// they are treated as variables defined at the top of the function.
//	These names should be used to document the meaning of the return values.
func subtract(x, y int) (res int) {
	res = x - y
	return
}

// more -> https://tour.golang.org/basics/7
