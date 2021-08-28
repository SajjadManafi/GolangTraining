package main

import "fmt"

func main() {

	a := "Hello"
	b := a
	b = "Hi"
	fmt.Println(a) // Hello -> not changed
	fmt.Println(b) // Hi

	x := "Hello"
	fmt.Printf("value: %v , Type: %T \n", x, x)   // value: Hello , Type: string
	fmt.Printf("value: %v , Type: %T \n", &x, &x) // value: 0xc00010a050 (for me) , Type: *string
	// & gives the address of x in memory

	var y *string = &x //or y := &x

	fmt.Printf("value: %v , Type: %T \n", y, y)   // value: 0xc00010a050 (for me) , Type: *string
	fmt.Printf("value: %v , Type: %T \n", *y, *y) // value: Hello , Type: string
	// * gives the value

	*y = "Hi"
	fmt.Println(x) // Hi -> changed

	// Pointers allow you to share a value stored in some memory location. Use pointers when
	// 1. you don’t want to pass around a lot of data
	// 2. you want to change the data at a location
	// "Everything in Go is pass by value"

	num := 10
	changerWithoutPointer(num)
	fmt.Println("Without Pointer ", num) // 10

	fmt.Println("")

	changerWithPointer(&num)
	fmt.Println("With Pointer ", num) // 210

	// more :
	//			https://tour.golang.org/moretypes/1
	//			method sets : https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver

}

func changerWithoutPointer(x int) {
	x += 200
	fmt.Println("in func changerWithoutPointer: ", x) // x + 200
}

func changerWithPointer(x *int) {
	*x += 200
	fmt.Println("in func changerWithPointer: ", *x) // x + 200
}
