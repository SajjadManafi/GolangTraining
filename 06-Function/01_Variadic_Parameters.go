package main

import (
	"fmt"
)

func main() {

	typePrint(2, 3, 4, 5, 6, 7) // [2 3 4 5 6 7] \n []int

	sum := sumFunc(90, 1, 1, 68, -7, 23)
	fmt.Println(sum) // 176

	numList := []int{900, 901, 902, 903, 904, 905}
	res := sumFunc(numList...)
	fmt.Println(res) // 5415

	sum2 := sumFunc()
	typePrint()       // [] \n []int
	fmt.Println(sum2) // 0

}

func typePrint(a ...int) {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

func sumFunc(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}
