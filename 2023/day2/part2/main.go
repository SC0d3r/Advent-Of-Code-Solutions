package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := ioutil.ReadFile("inp.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	total := 0
	start := time.Now()
	for _, l := range lines {
		_, hands := ParseLine(l)
		res := Gmp(hands)
		total += res[0] * res[1] * res[2]
	}
	elpd := time.Since(start)

	fmt.Println("total is", total, "elapsed is", elpd)
}

func Gmp(hands [][3]int) [3]int {
	res := [3]int{0, 0, 0}

	for _, hand := range hands {
		if hand[0] > res[0] {
			res[0] = hand[0]
		}

		if hand[1] > res[1] {
			res[1] = hand[1]
		}

		if hand[2] > res[2] {
			res[2] = hand[2]
		}
	}

	return res
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
