package main

import (
	"fmt"
	"runtime"
	"sync"
)

// What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?
// This concept is called mutual exclusion, and the conventional name for the data
// structure that provides it is mutex.

// We can define a block of code to be executed in mutual exclusion by surrounding it with
// a call to Lock and Unlock as shown on the Inc method.

// more :
//			https://pkg.go.dev/sync#Mutex.Unlock
//			https://tour.golang.org/concurrency/9

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//A locked Mutex is not associated with a particular goroutine.
			// It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
