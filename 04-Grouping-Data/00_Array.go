package main

import "fmt"

func main() {

	var x [4]int
	fmt.Println(x)

	//set a value at an index -> array[index] = value
	x[3] = 87
	fmt.Println(x)

	//get a value--> array[index].
	fmt.Println(x[2])

	// len returns the length of an array.
	fmt.Println(len(x))

	// using "RANGE"
	y := [4]float64{91.2, 23.3, 12.5, 4.1}

	for i, v := range y {
		fmt.Printf("index: %d , value: %.2f\n", i, v)
	}

	//two demical array:
	var t [2][2]int
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			t[i][j] = i * j
		}
	}

	fmt.Println(t)

	var alphabet [58]string

	for i := 65; i <= 122; i++ {
		alphabet[i-65] = string(rune(i))
	}

	fmt.Println(alphabet)
}
