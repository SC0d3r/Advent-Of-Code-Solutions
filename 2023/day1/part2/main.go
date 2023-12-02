package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var nMap = map[string]int{
	// "zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	file, err := ioutil.ReadFile("inp.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(file)

	lines := strings.Split(content, "\n")

	start := time.Now()

	total := 0
	for _, ln := range lines {
		n1, n2 := parse(ln)
		total += n1*10 + n2
	}

	elapsed := time.Since(start)

	fmt.Println("total is", total, "elapsed is", elapsed)
	// fmt.Println(content)
}

func parse(line string) (int, int) {
	res := make([]int, 0)

	ns := strings.Split(line, "")
	for i, n := range ns {
		x, err := strconv.Atoi(n)
		if err != nil {
			// n was not a valid number
			if i < 2 {
				continue
			}

			maybeNumber, err := ParseEngNumberRepr(ns, i)

			// fmt.Println("maybeNumber", maybeNumber)

			if err != nil {
				continue
			}

			res = append(res, maybeNumber)
		} else {
			// n was a valid number
			res = append(res, x)
		}
	}

	if len(res) == 0 {
		return 0, 0
	}

	return res[0], res[len(res)-1]
}

func ParseEngNumberRepr(ns []string, i int) (int, error) {
	y := i + 1

	for key, val := range nMap {
		if len(key) > y {
			continue
		}

		// fmt.Println("len(ns)", len(ns), "i is", i)

		if key == strings.Join(ns[y-len(key):y], "") {
			return val, nil
		}
	}

	return 0, fmt.Errorf("no eng number repr found")
}
