package main

import "fmt"

func main() {

	// this sruct doesnt any name...
	p := struct {
		name string
		age  int
	}{
		name: "Harry",
		age:  21,
	}

	fmt.Println(p)
}
