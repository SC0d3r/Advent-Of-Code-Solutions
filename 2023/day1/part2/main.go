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

		if err == nil {
			// n was a valid number
			res = append(res, x)
			continue
		}

		// n was not a valid number
		if i < 2 {
			continue
		}

		num, err := ParseEngNumberRepr(line, i)

		if err != nil {
			continue
		}

		res = append(res, num)
	}

	return res[0], res[len(res)-1]
}

func ParseEngNumberRepr(line string, i int) (int, error) {
	y := i + 1

	for key, val := range nMap {
		if len(key) > y {
			continue
		}

		if key == line[y-len(key):y] {
			return val, nil
		}
	}

	return 0, fmt.Errorf("no eng number repr found")
}
