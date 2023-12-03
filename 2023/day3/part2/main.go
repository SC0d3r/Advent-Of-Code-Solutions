package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("inp.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	// making the grid
	grid := make([][]interface{}, 0)

	for i, line := range lines {

		grid = append(grid, make([]interface{}, len(line)))

		chars := strings.Split(line, "")
		for j, char := range chars {
			x, err := strconv.Atoi(string(char))

			if err != nil {
				grid[i][j] = char
			} else {
				grid[i][j] = x
			}
		}
	}

	nMap := make(map[[2]int]int)
	singleAdjacentStarMap := make(map[[2]int]bool)

	// traverse the grid
	i := 0
	for i < len(grid) {
		j := 0
		for j < len(grid[i]) {
			if isNum(grid[i][j]) {
				res, specialChar, spI, spJ := adj(i, j, grid)

				if !res || specialChar != "*" {
					j++
					continue
				}

				newJ, wholeNumber := getWholeNumber(i, j, grid)

				key := [2]int{spI, spJ}

				if nMap[key] == 0 {
					nMap[key] = wholeNumber
					singleAdjacentStarMap[key] = true
				} else {
					nMap[key] *= wholeNumber
					singleAdjacentStarMap[key] = false
				}

				// jump the j after the whole number
				j = newJ
			} else {
				j++
			}
		}
		i++
	}

	total := 0

	for key, val := range nMap {
		if singleAdjacentStarMap[key] {
			continue
		}
		total += val
	}

	fmt.Println("total is", total)
}

func getWholeNumber(i int, j int, grid [][]interface{}) (int, int) {

	res := strconv.Itoa(grid[i][j].(int))

	// backward check
	for sj := j - 1; sj >= 0; sj-- {
		if num, ok := grid[i][sj].(int); ok {
			res = fmt.Sprintf("%d%s", num, res)
			continue
		}
		break
	}

	// forward check
	sj := j + 1
	for sj < len(grid[i]) {
		if num, ok := grid[i][sj].(int); ok {
			res = fmt.Sprintf("%s%d", res, num)
			sj++
			continue
		}
		break
	}

	n, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println(err)
		return sj, 0
	}

	return sj, n
}

func adj(i int, j int, grid [][]interface{}) (bool, string, int, int) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			xx := i + x
			yy := j + y

			// invalid index
			if xx < 0 || yy < 0 || xx >= len(grid) || yy >= len(grid[i]) {
				continue
			}

			el := grid[xx][yy]
			if isNum(el) || el == "." {
				// this is a number
				continue
			}

			// this is a special char
			// means that this number is adjacent to special char
			return true, el.(string), xx, yy
		}
	}

	return false, "", 0, 0
}

func isNum(el interface{}) bool {
	if _, ok := el.(int); ok {
		return true
	}
	return false
}
