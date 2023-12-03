package main

import (
	"testing"
)

func TestGmp(t *testing.T) {
	a := [][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	res := Gmp(a)

	if !eq(res, []int{7, 8, 9}) {
		t.Errorf("Expected %v, got %v", []int{7, 8, 9}, res)
	}
}

func eq(a [3]int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
