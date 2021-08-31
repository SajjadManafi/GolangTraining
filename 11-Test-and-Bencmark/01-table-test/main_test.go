package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	// slice of tests
	tests := []test{
		//{data , answer}
		{[]int{10, 9}, 19},
		{[]int{10, -10}, 0},
		{[]int{1, 0, -1}, 0},
		{[]int{950, 16}, 966},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
