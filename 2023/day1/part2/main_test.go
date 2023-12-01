package main

import (
	"strings"
	"testing"
)

func TestParseEngNumberRepr(t *testing.T) {
	const str1 = "two1nine"
	ns := strings.Split(str1, "")

	xs, err := ParseEngNumberRepr(ns, 1)

	if err == nil {
		t.Fatal("it should return err for ns, i which i is less than 2", "xs returned is", xs)
	}

	ys, err := ParseEngNumberRepr(ns, 2)

	if err != nil {
		t.Fatal("it should not return error", "and should return 2", "but returned", ys)
	}

	ls, _ := ParseEngNumberRepr(ns, 7)
	if ls != 9 {
		t.Fatal("it should return 9", "but returned", ls)
	}
}
