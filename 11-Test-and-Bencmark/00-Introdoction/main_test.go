package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(90, 42)
	if x != 132 {
		t.Error("Expected", 132, "Got ", x)
	}
}
