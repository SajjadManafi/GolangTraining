package main

import "fmt"

func main() {

	s := "Hello World"
	fmt.Printf(s)
	fmt.Printf("%T\n", s)

	// slice of string characters ascii code
	byteS := []byte(s)
	fmt.Println(byteS)
	fmt.Printf("%T\n", byteS)

	// print Unicode Character of characters
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x -> %q \n", i, v, v)
	}

	// for More -> https://golang.org/ref/spec#String_types
	// and https://pkg.go.dev/fmt

}
