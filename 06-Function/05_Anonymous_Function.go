package main

import "fmt"

func main() {

	func(name string) {
		fmt.Println("Hello", name)
	}("Harry") // Harry

	func() {
		fmt.Println("Anonymous Func")
	}() // Anonymous Func

	x := func() int {
		return 2
	}()

	fmt.Println(x) // 2

	y := func() int {
		return 5
	}
	fmt.Printf("%T\n", y) // func() int

	z := y()
	fmt.Println(z) // 5

}
