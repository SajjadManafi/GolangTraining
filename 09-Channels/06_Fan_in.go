package main

import (
	"fmt"
	"sync"
)

//Taking values from many channels, and putting those values onto one channel.

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sendd(even, odd)

	go receivee(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func sendd(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receivee(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

// code source : Todd_McLeod
// https://play.golang.org/p/_CyyXQBCHe
// more : https://play.golang.org/p/buy30qw5MM
// more : https://play.golang.org/p/RzR3Kjrx7q
