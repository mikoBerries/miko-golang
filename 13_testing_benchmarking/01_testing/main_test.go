//test file for main file function
package main

import "testing"

//go test
func TestMathadding(t *testing.T) {
	expectedVal1 := 300
	expectedVal2 := 500
	x := Mathadding(200, 100)
	if x != 300 {
		t.Error("Expected", expectedVal1, "got", x)
	} else {
		t.Log("pass", x)
	}
	x = Mathadding(500, -100)
	if x != 400 {
		t.Error("Expected", expectedVal2, "got", x)
	}

	//table test multi testing value using struct

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{10, 20}, 30},
		test{[]int{100, 200}, 300},
		test{[]int{1, 2}, 3},
		test{[]int{0, 0}, 0},
	}
	for _, v := range tests {
		tempValue := Mathadding(v.data...)
		if tempValue != v.answer {
			t.Error("Expected", v.answer, "got", tempValue)
		}
	}

}
