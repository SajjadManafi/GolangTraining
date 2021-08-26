package main

import "fmt"

func main() {

	// for init; condition; post {} -> Like a C for

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	// Like a C while
	x := 1
	for x < 20 {
		x *= 2
		fmt.Println(x)
	}

	fmt.Println("-----------")

	// Like a C for(;;)
	y := 0
	for {

		fmt.Println(y)

		if y > 5 {
			fmt.Println("END!")
			break
		}

		y++
	}

	fmt.Println("-----------")

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("-----------")

	// Example -> Printing Ascii
	for j := 33; j <= 122; j++ {
		fmt.Printf("%v\t%#x\t%#U\n", j, j, j)
	}

	fmt.Println("-----------")

	s := "Good Morning!"

	for q, w := range s {
		fmt.Printf("at index position %d we have %q \n", q, w)
	}

	// for more -> https://golang.org/ref/spec#For_statements

}
