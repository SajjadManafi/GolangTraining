package main

import "fmt"

func main() {

	{
		// send-only channel
		c := make(chan<- int, 1)

		c <- 10

		fmt.Printf("%T\n", c)
		// we cant do this  (<-c)
	}

	{
		// recive-only channel

		c := make(<-chan int, 1)
		// we cant do this (c <- 10)

		// we can just use (<-c)

		fmt.Printf("%T\n", c)
	}

	{
		fmt.Println("*********")

		c := make(chan int)
		cr := make(<-chan int) // receive
		cs := make(chan<- int) // send
		fmt.Printf("c:%T\n", c)
		fmt.Printf("cr:%T\n", cr)
		fmt.Printf("cs:%T\n", cs)

		// specific to general doesn't assign
		//c = cr
		//c = cs

		// specific to specific doesn't assign
		//cs = cr

		// general to specific assigns
		cr = c
		cs = c

	}
	{
		fmt.Println("--------")

		c := make(chan int)
		cr := make(<-chan int) // receive
		cs := make(chan<- int) // send
		fmt.Printf("c:%T\n", c)
		fmt.Printf("cr:%T\n", cr)
		fmt.Printf("cs:%T\n", cs)

		// general to specific converts
		fmt.Println("-----")
		fmt.Printf("c\t%T\n", (<-chan int)(c))
		fmt.Printf("c\t%T\n", (chan<- int)(c))

		// specific to general doesn't convert
		//fmt.Printf("c\t%T\n", (chan int)(cr))
		//fmt.Printf("c\t%T\n", (chan int)(cs))

	}
}
