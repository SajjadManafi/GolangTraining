package main

import "fmt"

func main() {

	// EX 1:
	fmt.Println(calculation(2, 5, add))      // 7
	fmt.Println(calculation(9, 7, multiply)) // 63

	// EX 2:
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	all := SumFunc(l...)
	fmt.Println("all: ", all) //55

	ev := evenSum(SumFunc, l...)
	fmt.Println("even: ", ev) // 30
}

// EX 1:
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func calculation(x, y int, f func(a, b int) int) int {
	return f(x, y)
}

// EX 2:
func SumFunc(x ...int) int {
	sum := 0
	for _, num := range x {
		sum += num
	}

	return sum
}

func evenSum(f func(x ...int) int, y ...int) int {

	li := make([]int, 0, len(y))
	for _, num := range y {
		if num%2 == 0 {
			li = append(li, num)
		}
	}

	sum := f(li...)
	return sum

}
