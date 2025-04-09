package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

	now := time.Now()
	// all the Xs indexes
	// var wg sync.WaitGroup
	for i, row := range twoD {
		for j, char := range row {
			if char == "A" {
				if key := findXMas(i, j, twoD); key != "" && notInFoundMap(key, foundMap) {
					foundMap[key] = true
					founds++
				}

				// wg.Add(1)
				// go func() {
				// 	key := findXMas(i, j, twoD)
				// 	if key != "" {
				// 		foundMap[key] = true
				// 		founds++
				// 	}
				// 	wg.Done()
				// }()
			}
		}
	}
	// wg.Wait()

	log.Println("foundMap", foundMap)
	log.Println("elapsed", time.Now().Sub(now))
	log.Println("founds", founds)

	// for key := range foundMap {
	// 	log.Println("foundMap key", string(key))
	// }
}

func findXMas(i, j int, twoD [][]string) string {
	diag := findMAS(i, j, twoD, [][]int{{-1, -1}, {1, 1}})
	log.Println("diag is", diag)
	if diag == "" {
		return ""
	}

	revDiag := findMAS(i, j, twoD, [][]int{{-1, 1}, {1, -1}})
	log.Println("revDiag is", revDiag)
	if revDiag == "" {
		return ""
	}

	// combined key
	return diag + revDiag
}

func findMAS(i, j int, twoD [][]string, idx [][]int) string {
	if len(idx) != 2 {
		log.Fatal("this shouldnt happen")
	}

	log.Println("len(twoD)", len(twoD), "len(twoD[0])", len(twoD[0]), "i", i, "j", j)

	res := "A"
	key := "" + strconv.Itoa(i) + strconv.Itoa(j)
	for _, id := range idx {
		key += strconv.Itoa(i+id[0]) + strconv.Itoa(j+id[1])

		// check invalid indexes
		if i+id[0] < 0 || (i+id[0] >= len(twoD)) || j+id[1] < 0 || j+id[1] >= len(twoD[0]) {
			continue
		}

		log.Println("checking i, j", i+id[0], j+id[1], "char", twoD[i+id[0]][j+id[1]])

		if len(res) == 1 {
			// prepend
			res = twoD[i+id[0]][j+id[1]] + res
		} else {
			// append
			res += twoD[i+id[0]][j+id[1]]
		}
	}

	if res != "MAS" && res != "SAM" {
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
