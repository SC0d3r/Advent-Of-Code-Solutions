package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const EMPTY = "."

type Step struct {
	val   uint64
	empty bool
	id    uint64
}

// applied steps
type AP struct {
	val   uint64
	empty bool
}

func main() {
	file, err := os.Open("./2024/day9/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	steps := generateSteps(string(data))
	log.Println("steps", steps)

	res := applySteps(steps)

	log.Println("after apply res is", res)

	reorderStep(res)

	log.Println("after reorder res is", res)

	var checksum uint64 = 0
	for i, s := range res {
		checksum += uint64(i) * s.val
	}
	log.Println("checksum", checksum)
}

func reorderStep(aps []AP) {
	for {
		left := findLeftMostEmptyIdx(aps)
		right := findRightmostNonEmptyIdx(aps)
		// Stop if the leftmost free space is not to the left of any file block.
		if left >= right {
			break
		}
		// Move the block: set the leftmost free position to the file block at "right",
		// and mark the old position as free.
		aps[left] = aps[right]
		aps[right] = AP{val: 0, empty: true}
	}
}

func findLeftMostEmptyIdx(aps []AP) int {
	for i, n := range aps {
		if n.empty {
			return i
		}
	}
	return len(aps) // should not happen if there is at least one empty block
}

func applySteps(steps []Step) []AP {
	res := []AP{}
	for _, s := range steps {
		if s.empty {
			// res += repeatStr(EMPTY, s.val)
			res = append(res, repeat(true, 0, s.val)...)
		} else {
			// file
			res = append(res, repeat(false, s.id, s.val)...)
			// res += repeatNum(s.id, s.val)
		}
	}

	return res
}

func repeat(isEmpty bool, id uint64, howMany uint64) []AP {
	res := []AP{}
	var i uint64 = 0
	for ; i < howMany; i = i + 1 {
		res = append(res, AP{val: id, empty: isEmpty})
	}

	return res
}

func generateSteps(num string) []Step {
	steps := []Step{}

	strs := strings.Split(num, "")
	nums := []int{}
	for _, s := range strs {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, x)
	}

	isFile := false
	id := 0
	for i, x := range nums {
		isFile = i%2 == 0
		if isFile {
			steps = append(steps, Step{
				val:   uint64(x),
				id:    uint64(id),
				empty: false,
			})
			id++
		} else {
			steps = append(steps, Step{
				val:   uint64(x),
				id:    uint64(id),
				empty: true,
			})
		}
	}

	return steps
}

func findRightmostNonEmptyIdx(aps []AP) int {
	for i := len(aps) - 1; i >= 0; i-- {
		if !aps[i].empty {
			return i
		}
	}
	return -1 // should never happen if there is at least one file block
}
