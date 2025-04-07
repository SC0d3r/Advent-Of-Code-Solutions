package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
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
	right := make(map[int]int, 0)

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

		left = append(left, leftInt)
		right[rightInt] = right[rightInt] + 1
	}

	fmt.Println("left", left)
	fmt.Println("right", right)

	diffs := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		diffs[i] = left[i] * int(math.Max(0, float64(right[left[i]])))
	}

	fmt.Println("diffs", diffs)

	sum := 0
	for _, x := range diffs {
		sum += x
	}

	fmt.Println("sum is", sum)
}
