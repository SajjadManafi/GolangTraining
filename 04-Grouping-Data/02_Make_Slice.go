package main

import "fmt"

func main() {

	// make (type(Pointer), len , Capacity)
	x := make([]int, 0, 0)
	fmt.Println(x)
	fmt.Printf("len: %d , cap: %d \n", len(x), cap(x))

	x = append(x, 1)
	fmt.Println(x)
	fmt.Printf("len: %d , cap: %d \n", len(x), cap(x))

	x = append(x, 5)
	fmt.Println(x)
	fmt.Printf("len: %d , cap: %d \n", len(x), cap(x))

	x = append(x, 7)
	fmt.Println(x)
	fmt.Printf("len: %d , cap: %d \n", len(x), cap(x))

	// Each time you append to Slice if you exceed the number of Caps, Behind the scenes
	// Go runner creates a value array twice the size of the previous Cap

	number := make([]int, 0, 10)

	for i := 0; i <= 20; i++ {
		number = append(number, i)
		fmt.Printf("i:%d , len:%d , cap:%d\n", i, len(number), cap(number))
	}

}
