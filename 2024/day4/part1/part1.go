package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2024/day4/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	twoD := make([][]string, 0)

	for _, l := range lines {
		log.Println("lines", l)
	}

	for _, line := range lines {
		chars := strings.Split(strings.TrimSpace(line), "")
		twoD = append(twoD, chars)
	}

	foundMap := make(map[string]bool, 0)

	founds := 0

	checks := []func(int, int, [][]string) string{
		linearForward,
		linearBackward,
		linearUpward,
		linearDownward,
		linearDiagonalNorthEast,
		linearDiagonalNorthWest,
		linearDiagonalSouthWest,
		linearDiagonalSouthEast,
	}

	// all the Xs indexes
	for i, row := range twoD {
		for j, char := range row {
			if char == "X" {
				for _, check := range checks {
					if key := check(i, j, twoD); key != "" && notInFoundMap(key, foundMap) {
						foundMap[key] = true
						founds++
					}
				}
			}
		}
	}

	log.Println("founds", founds)
	log.Println("foundMap", foundMap)

	// for key := range foundMap {
	// 	log.Println("foundMap key", string(key))
	// }
}

func linearForward(i, j int, twoD [][]string) string {
	// log.Println("checking linearForward", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}})
}

func linearBackward(i, j int, twoD [][]string) string {
	// log.Println("checking linearBackward", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{0, -1}, []int{0, -2}, []int{0, -3}})
}

func linearUpward(i, j int, twoD [][]string) string {
	// log.Println("checking linearUpwar", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{-1, 0}, []int{-2, 0}, []int{-3, 0}})
}

func linearDownward(i, j int, twoD [][]string) string {
	// log.Println("checking linearDownward", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 0}})
}

func linearDiagonalSouthEast(i, j int, twoD [][]string) string {
	// log.Println("checking linearDiag", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{1, 1}, []int{2, 2}, []int{3, 3}})
}

func linearDiagonalSouthWest(i, j int, twoD [][]string) string {
	// log.Println("checking linearRevDiag", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{1, -1}, []int{2, -2}, []int{3, -3}})
}

func linearDiagonalNorthEast(i, j int, twoD [][]string) string {
	return findXMAS(i, j, twoD, [][]int{[]int{-1, 1}, []int{-2, 2}, []int{-3, 3}})
}

func linearDiagonalNorthWest(i, j int, twoD [][]string) string {
	// log.Println("checking linearRevDiag", i, j)
	return findXMAS(i, j, twoD, [][]int{[]int{-1, -1}, []int{-2, -2}, []int{-3, -3}})
}

func findXMAS(i, j int, twoD [][]string, idx [][]int) string {
	if len(idx) != 3 {
		return ""
	}

	log.Println("len(twoD)", len(twoD), "len(twoD[0])", len(twoD[0]), "i", i, "j", j)

	res := "X"
	key := "" + strconv.Itoa(i) + strconv.Itoa(j)
	for _, id := range idx {
		key += strconv.Itoa(i+id[0]) + strconv.Itoa(j+id[1])

		// check invalid indexes
		if i+id[0] < 0 || (i+id[0] >= len(twoD)) || j+id[1] < 0 || j+id[1] >= len(twoD[0]) {
			continue
		}

		log.Println("checking i, j", i+id[0], j+id[1], "char", twoD[i+id[0]][j+id[1]])
		res += twoD[i+id[0]][j+id[1]]
	}

	if res != "XMAS" {
		return ""
	}

	// here we found it
	return key
}

func notInFoundMap(key string, foundMap map[string]bool) bool {
	if _, ok := foundMap[key]; ok {
		return false
	}

	return true
}
