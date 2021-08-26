package main

import (
	"fmt"
	"time"
)

func main() {

	switch {
	case 1 == 1:
		fmt.Println("1 == 1")
		// switch break here. because case(1 == 1) is true
	case 2 < 6:
		fmt.Println("2 < 6")
	}

	// use fallthrough
	switch {
	case 1 == 1:
		fmt.Println("1 == 1 -> true")
		fallthrough
		// switch not breaking here
	case 2 < 6:
		fmt.Println("2 < 6")
	}

	// switch with default
	switch {
	case true:
		fmt.Println("true")
		fallthrough
	case 9 < 3:
		fmt.Println("9 < 3")
		fallthrough
	case len("Hey") == 3:
		fmt.Println(`len of "Hey" is 3`)
		fallthrough
	default:
		fmt.Println("default")
		//If none of the cases were true. Or we did use fallthrough. It runs
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case 0 > 10:
		fmt.Println("this should not print")
	default:
		fmt.Println("this is default")
	}

	message := "Hey"
	switch message {

	case "Hello", "Hey", "Hi":
		fmt.Println("Hi")

	case "How are you?":
		fmt.Println("Thanks. im fine.")
	default:
		fmt.Println("what?!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

}
