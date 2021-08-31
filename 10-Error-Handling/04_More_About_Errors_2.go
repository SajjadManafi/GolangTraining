package main

import (
	"fmt"
	"time"
)

// more : https://tour.golang.org/methods/19

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
