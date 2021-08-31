package main

import "fmt"

func main() {
	fmt.Println(mySum(12, 17, 901))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
