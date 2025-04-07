package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)

	file, err := os.Open("./2024/day1/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// data, err := io.ReadAll("./2024/day1/inp1.txt")
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		sep := strings.Split(strings.TrimSpace(line), "  ")

		leftInt, err := strconv.Atoi(strings.TrimSpace(sep[0]))
		if err != nil {
			log.Fatal(err)
		}

		rightInt, err := strconv.Atoi(strings.TrimSpace(sep[1]))
		if err != nil {
			log.Fatal(err)
		}

		left = InsertSorted(left, leftInt)
		right = InsertSorted(right, rightInt)
	}

	fmt.Println("left", left)
	fmt.Println("right", right)

	diffs := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		diffs[i] = int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	fmt.Println("diffs", diffs)

	sum := 0
	for _, x := range diffs {
		sum += x
	}

	fmt.Println("sum is", sum)

	// for i, l := range data {
	// 	fmt.Println("line", i, "content", l)
	// }

	// fmt.Println(string(data))
}

func InsertSorted(s []int, e int) []int {
	// Find the insertion point using binary search
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })

	// Append a placeholder value at the end
	s = append(s, 0)

	// fmt.Println("i is", i, "e is ", e, "s is", s)
	// log.Fatal("here")

	// Shift elements to the right, starting from index `i`
	copy(s[i+1:], s[i:])

	s[i] = e

	return s
}
