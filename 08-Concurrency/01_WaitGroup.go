package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.
// Writing concurrent code is super easy: all we do is put “go” in front of a function or method call.

// for more:
//				https://pkg.go.dev/sync
//				https://golang.org/doc/effective_go#concurrency

func main() {

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t ", runtime.NumGoroutine())

	// lunch another Goroutine with "go"
	// if we haven't WaitGroup -> our control flow doesn't have  to wait for this  Goroutine to execute
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t ", runtime.NumGoroutine())

	wg.Wait()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i, " foo")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i, " bar")
	}
}
