package main

import "fmt"

var x int

func main() {

	{
		var y int
		fmt.Println("y: ", y)
	}
	// we cant use y here -> fmt.Println(y)

	hey()

	a := incrementor()
	b := incrementor()

	fmt.Println("a: ", a()) // 1
	fmt.Println("a: ", a()) // 2
	fmt.Println("a: ", a()) // 3
	fmt.Println("a: ", a()) // 4

	fmt.Println("b: ", b()) // 1
	fmt.Println("b: ", b()) // 2
	fmt.Println("b: ", b()) // 3

}

func hey() {
	// but we can use x here
	x++
	fmt.Println("x in hey: ", x)
}

func incrementor() func() int {
	var z int
	return func() int {
		z++
		return z
	}
}
