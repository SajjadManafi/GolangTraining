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
