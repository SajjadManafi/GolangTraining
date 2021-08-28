package main

import "fmt"

func main() {

	a := func() {
		fmt.Println("Heeeeeey!")
	}

	a()

	addOne := func(x int) int {
		return x + 1
	}

	oldNum := 20
	newNum := addOne(oldNum)
	fmt.Println(newNum)
}
