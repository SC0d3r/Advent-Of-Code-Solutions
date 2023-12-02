package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	const line string = "Game 1: 1 green, 2 blue; 15 blue, 12 red, 2 green; 4 red, 6 blue; 10 blue, 8 red; 3 red, 12 blue; 1 green, 12 red, 8 blue"

	id, res := ParseLine(line)

	if res == nil {
		t.Error("ParseLine failed", "id is", id)
	}

	if id != 1 {
		t.Error("ParseLine failed", "id should be 1 but instead got", id)
	}

	if !equal(res[0], []int{0, 1, 2}) {
		t.Error("ParseLine failed", "res[0] should be", []int{0, 1, 2}, "but got", res[0])
	}

	if !equal(res[1], []int{12, 2, 15}) {
		t.Error("ParseLine failed", "res[1] should be", []int{12, 2, 15}, "but got", res[1])
	}

	if !equal(res[2], []int{4, 0, 6}) {
		t.Error("ParseLine failed", "res[2] should be", []int{4, 0, 6}, "but got", res[2])
	}

	if !equal(res[3], []int{8, 0, 10}) {
		t.Error("ParseLine failed", "res[3] should be", []int{8, 0, 10}, "but got", res[3])
	}

	if !equal(res[4], []int{3, 0, 12}) {
		t.Error("ParseLine failed", "res[4] should be", []int{3, 0, 12}, "but got", res[4])
	}

	if !equal(res[5], []int{12, 1, 8}) {
		t.Error("ParseLine failed", "res[5] should be", []int{12, 1, 8}, "but got", res[5])
	}

}

func equal(a [3]int, b []int) bool {
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
