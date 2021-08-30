package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// more :
//			https://pkg.go.dev/sync/atomic
//			https://gobyexample.com/atomic-counters

// It’s safe to access counter now because we know no other goroutine is writing to it.

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("count:\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
