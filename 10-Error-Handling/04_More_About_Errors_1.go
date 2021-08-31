package main

import (
	"errors"
	"fmt"
)

// more : https://gobyexample.com/errors

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

func main() {

	fmt.Println("The number is 23")
	one, err := f1(23)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(one)
	}

	fmt.Println("The number is 42")
	two, err := f1(42)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(two)
	}
}
