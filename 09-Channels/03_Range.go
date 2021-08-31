package main

import (
	"fmt"
)

// more:
//		 https://gobyexample.com/closing-channels
//		 https://gobyexample.com/range-over-channels

func main() {

	c := make(chan int)

	go sendC(c)

	receiveC(c)

}

func sendC(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func receiveC(c <-chan int) {
	for v := range c {
		fmt.Println("v:", v)
	}
}
