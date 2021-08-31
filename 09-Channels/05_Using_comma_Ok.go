package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c

	fmt.Println(v, ok) // 42 true

	v, ok = <-c

	fmt.Println(v, ok) // 0 false
}
