package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	// "time"
)

func main() {
	file, err := ioutil.ReadFile("inp.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	actual := []int{12, 13, 14}

	total := 0
	for _, l := range lines {
		id, hands := ParseLine(l)
		if isValid(hands, actual) {
			total += id
		}
	}

	fmt.Println("total is", total)
}

func isValid(hands [][3]int, actual []int) bool {
	for _, cubes := range hands {
		if cubes[0] > actual[0] || cubes[1] > actual[1] || cubes[2] > actual[2] {
			return false
		}
	}

	return true
}

func ParseLine(line string) (int, [][3]int) {
	xs := strings.Split(line, ":")
	idPart, cubesPart := xs[0], strings.TrimSpace(xs[1])

	res := make([][3]int, 0)

	id, err := strconv.Atoi(strings.Split(idPart, " ")[1])

	if err != nil {
		fmt.Println(err)
		return -1, nil
	}

	ys := strings.Split(cubesPart, ";")
	for i, y := range ys {
		// fmt.Println("i is", i, "y is", y)
		res = append(res, [3]int{})

		zs := strings.Split(y, ",")
		for _, z := range zs {
			// count, color
			cc := strings.Split(strings.TrimSpace(z), " ")
			count, err := strconv.Atoi(cc[0])
			color := cc[1]

			if err != nil {
				fmt.Println(err)
				return -1, nil
			}

			if color == "red" {
				res[i][0] = count
			}

			if color == "green" {
				res[i][1] = count
			}

			if color == "blue" {
				res[i][2] = count
			}
		}
	}

	return id, res
}
