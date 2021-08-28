package main

import "fmt"

func main() {

	// map[type]type
	x := map[string]int{
		"One": 1,
		"Two": 2,
	}

	fmt.Println(x)        // map[One:1 Two:2]
	fmt.Println(x["One"]) // 1

	val, exist := x["Five"]
	fmt.Printf("value: %d , exist: %v\n", val, exist) // value: 0 , exist: false

	// add element

	x["Three"] = 3
	fmt.Println(x) // map[One:1 Three:3 Two:2]

	// using range
	for key, val := range x {
		fmt.Printf("%s -> %d\n", key, val)
	}

	// delete
	delete(x, "Three")
	fmt.Println(x) // map[One:1 Two:2]

	// Nothing will happen.
	delete(x, "Five")
	fmt.Println(x) // map[One:1 Two:2]

	// make map with make(map[string]int)
	m := make(map[int]string)
	m[1] = "One"
	m[2] = "Two"
	fmt.Println(m) //	map[1:One 2:Two]
}
