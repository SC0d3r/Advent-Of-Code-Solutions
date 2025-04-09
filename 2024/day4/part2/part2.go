package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	file, err := os.Open("./2024/day4/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	twoD := make([][]string, 0)
	for _, line := range lines {
		chars := strings.Split(strings.TrimSpace(line), "")
		twoD = append(twoD, chars)
	}

	results := make(chan string)
	done := make(chan bool)
	foundMap := make(map[string]bool)
	founds := 0

	now := time.Now()

	// Start a collector goroutine to update the foundMap and counter.
	go func() {
		for key := range results {
			if !foundMap[key] {
				foundMap[key] = true
				founds++
			}
		}
		done <- true
	}()

	var wg sync.WaitGroup

	// Launch goroutines for each relevant cell concurrently.
	for i, row := range twoD {
		for j, char := range row {
			if char == "A" {
				wg.Add(1)
				go func() {
					defer wg.Done()
					key := findXMas(i, j, twoD)
					if key != "" {
						results <- key
					}
				}()
			}
		}
	}

	// Wait for all goroutines to finish then close the results channel.
	wg.Wait()
	close(results)
	// Wait for the collector to finish processing.
	<-done

	// log.Println("foundMap", foundMap)
	log.Println("elapsed", time.Now().Sub(now))
	log.Println("founds", founds)
}

func findXMas(i, j int, twoD [][]string) string {
	diag := findMAS(i, j, twoD, [][]int{{-1, -1}, {1, 1}})
	if diag == "" {
		return ""
	}
	revDiag := findMAS(i, j, twoD, [][]int{{-1, 1}, {1, -1}})
	if revDiag == "" {
		return ""
	}
	return diag + revDiag
}

func findMAS(i, j int, twoD [][]string, idx [][]int) string {
	if len(idx) != 2 {
		log.Fatal("this shouldn't happen")
	}

	res := "A"
	key := "" + strconv.Itoa(i) + strconv.Itoa(j)
	for _, id := range idx {
		key += strconv.Itoa(i+id[0]) + strconv.Itoa(j+id[1])
		if i+id[0] < 0 || (i+id[0] >= len(twoD)) || j+id[1] < 0 || j+id[1] >= len(twoD[0]) {
			continue
		}
		if len(res) == 1 {
			res = twoD[i+id[0]][j+id[1]] + res
		} else {
			res += twoD[i+id[0]][j+id[1]]
		}
	}
	if res != "MAS" && res != "SAM" {
		return ""
	}
	return key
}
