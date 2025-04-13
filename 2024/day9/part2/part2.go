package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const EMPTY = "."

type Space struct {
	idx     int
	howMany int
}

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

	print(res)

	fitSpace(res)

	log.Println("after reorder res is", res)
	print(res)

	var checksum uint64 = 0
	for i, s := range res {
		checksum += uint64(i) * s.val
	}
	log.Println("checksum", checksum)
}

func fitSpace(aps []AP) {
	rIdx := len(aps) - 1
	// lIdx := 0
	// processedAps := map[uint64]struct{}{}

	for {
		_, howMany, movedRIdx := biggestRightMostAP(rIdx, aps)
		// log.Println("moving rIdex", movedRIdx, "prev", rIdx)
		rIdx = movedRIdx
		howManySpaces, lIdx := findLeftMostEmptySpace(aps, howMany)

		// _, exists := processedAps[ap.val]
		// log.Println("exists", exists, ap)
		// if exists {
		// 	break
		// }

		// processedAps[ap.val] = struct{}{}
		if rIdx <= 0 {
			break
		}

		// log.Println("biggest howMany", howMany, ap, "most left", howManySpaces, "lIdx", lIdx, "rIdx", rIdx)

		if howManySpaces < howMany {
			// cannot fit
			// so we have to skip these aps
			rIdx -= howMany
			// log.Println("adjusting rIdx", rIdx)
			continue
		}

		if lIdx >= rIdx {
			// finish?
			// break
			rIdx -= howMany
			continue
		}

		// here means it can fit
		// print(aps)
		// log.Println("21", aps[21])
		for i := 0; i < howMany; i++ {
			// log.Println("placing", aps[rIdx], "into", aps[lIdx])
			aps[lIdx] = aps[rIdx]
			aps[rIdx] = AP{val: 0, empty: true}
			lIdx++
			rIdx--
		}

		// rIdx -= howMany
	}
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

func print(aps []AP) {
	res := ""
	for _, ap := range aps {
		if ap.empty {
			res += EMPTY
		} else {
			res += strconv.Itoa(int(ap.val))
		}
	}

	fmt.Println(res)
}

func biggestRightMostAP(rIdx int, aps []AP) (ap AP, howMany int, movedRIdx int) {
	found := false
	for i := rIdx; i >= 0; i-- {
		if found {
			if ap.val != aps[i].val {
				return
			}
			howMany++
			continue
		}

		if aps[i].empty {
			continue
		}
		found = true
		movedRIdx = i
		ap = aps[i]
		howMany = 1
	}

	return
}

func findLeftMostEmptySpace(aps []AP, neededHowMany int) (howMany int, lIdx int) {
	spaces := []Space{}
	lIdx = -1

	// found := false
	for i := 0; i < len(aps); i++ {
		if aps[i].empty {
			// found = true
			if lIdx == -1 {
				lIdx = i
			}

			howMany++
		} else {
			if howMany > 0 {
				spaces = append(spaces, Space{
					idx:     lIdx,
					howMany: howMany,
				})
			}

			howMany = 0
			lIdx = -1
			// found = false
		}
	}

	// find the biggest
	biggestSpace := spaces[0]
	for _, sp := range spaces {
		if sp.howMany >= neededHowMany {
			return sp.howMany, sp.idx
		}

		if biggestSpace.howMany < sp.howMany {
			biggestSpace = sp
		}
	}

	return biggestSpace.howMany, biggestSpace.idx
}
