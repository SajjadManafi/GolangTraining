package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Race Condition Problem
// Race conditions are not good code. A race condition will give unpredictable results.

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter

			//Gosched yields the processor, allowing other goroutines to run.
			// It does not suspend the current goroutine, so execution resumes automatically.
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
