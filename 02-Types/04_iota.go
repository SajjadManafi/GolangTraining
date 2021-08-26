package main

import "fmt"

const (
	q = iota
	w = iota
	e = iota
)

const (
	r = iota
	t
	u
)

func main() {

	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(e)
	// 0 , 1 , 2

	fmt.Println(r)
	fmt.Println(t)
	fmt.Println(u)

	//automatically increment
	//it's only when you have const word again that it resets.

}
