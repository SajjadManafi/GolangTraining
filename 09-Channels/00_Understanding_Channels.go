package main

import (
	"fmt"
)

//channels block
//	they are synchronized
//	they have to pass/receive the value at the same time
//	just like runners in a relay race have to pass / receive the baton to each other at the same time
//	one runner can’t pass the baton at one moment
//	and then, later, have the other runner receive the baton
//	the baton is passed/received by the runners at the same time
//	the value is passed/received synchronously; at the same time
//	channels allow us to pass values between goroutines

func main() {
	{
		// this block doesnt run
		// make channel
		// c := make(chan int)

		//putting values on a channel (c <- 12)
		// c <- 12

		//taking values off of a channel
		// fmt.Println(<-c)
	}

	{
		c := make(chan int)

		go func() {
			c <- 12
		}()

		fmt.Println(<-c)
	}

	// or use buffered channels
	{
		c := make(chan int, 1)

		c <- 80
		// we cant do this again (c <- value)

		fmt.Println(<-c)
	}

	{
		c := make(chan int, 2)

		c <- 90
		c <- 56

		fmt.Println(<-c) // 90
		fmt.Println(<-c) // 56
	}
	// or
	{
		c := make(chan int, 1)

		c <- 102

		go func() {
			c <- 63
		}()

		fmt.Println(<-c) // 102
		fmt.Println(<-c) // 63
	}

}
