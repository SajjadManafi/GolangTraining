package main

import "fmt"

func main() {

	c := make(chan int)

	go sender(c, 92)

	receiver(c)
}

func sender(c chan<- int, number int) {
	c <- number
}

func receiver(c <-chan int) {
	fmt.Println(<-c)
}
