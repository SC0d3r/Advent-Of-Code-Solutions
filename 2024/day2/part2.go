package main

import (
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2024/day2/inp1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	safeReports := make([][]string, 0)

	for _, line := range lines {
		levels := strings.Split(strings.TrimSpace(line), " ")
		if isReportSafeWithTolerate(levels) {
			safeReports = append(safeReports, levels)
		}
	}

	log.Println("safe reports are", len(safeReports))
}

func isReportSafeWithTolerate(report []string) bool {
	for i := 0; i < len(report); i++ {
		if isReportSafe(without(report, i)) {
			return true
		}
	}

	return false
}

func without(arr []string, index int) []string {
	ret := make([]string, 0)

	for i, v := range arr {
		if i == index {
			continue
		}

		ret = append(ret, v)
	}

	return ret
}

func isReportSafe(report []string) bool {
	isIncreasing := false
	isDecreasing := false

	rl := len(report)
	for i := 0; i < rl-1; i++ {
		lvl, err := strconv.Atoi(report[i])
		if err != nil {
			log.Fatal(err)
		}
		nextLvl, err := strconv.Atoi(report[i+1])
		if err != nil {
			log.Fatal(err)
		}

		if lvl == nextLvl {
			return false
		}

		if lvl < nextLvl && isIncreasing {
			return false
		}

		if lvl > nextLvl && isDecreasing {
			return false
		}

		if lvl < nextLvl {
			isDecreasing = true
		} else {
			isIncreasing = true
		}

		diff := math.Abs(float64(lvl - nextLvl))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
