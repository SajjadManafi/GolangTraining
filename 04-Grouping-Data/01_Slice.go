package main

import "fmt"

func main() {

	x := []int{9, 5, 35, 26, 873, 4}
	fmt.Println(x) //[9 5 35 26 873 4]

	fmt.Println(x[1:]) // from index 1 to end -> [5 35 26 873 4]

	fmt.Println(x[1:4]) // from index 1 to 4 -> [5 35 26]

	// append to slice
	x = append(x, 46, 12, 2021)
	fmt.Println(x) // [9 5 35 26 873 4 46 12 2021]

	// append slice to slice
	y := []int{2022, 2023, 2024, 2025}

	x = append(x, y...)
	fmt.Println(x) // [9 5 35 26 873 4 46 12 2021 2022 2023 2024 2025]

	// deleting from slice

	x = append(x[:2], x[4:]...)
	fmt.Println(x) // [9 5 873 4 46 12 2021 2022 2023 2024 2025] -> 35 and 25 deleted

	// two demical slice
	z := [][]int{x, y}
	fmt.Println(z) // [[9 5 873 4 46 12 2021 2022 2023 2024 2025] [2022 2023 2024 2025]]
}
