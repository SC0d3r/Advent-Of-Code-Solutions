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

	content := string(file)

	lines := strings.Split(content, "\n")

	total := 0
	for _, ln := range lines {
		n1, n2 := getNumbers(ln)
		total += n1*10 + n2
	}

	fmt.Println("total is", total)
	// fmt.Println(content)
}

func getNumbers(str string) (int, int) {
	res := make([]int, 0)

	xs := strings.Split(str, "")
	for _, x := range xs {
		xInt, err := strconv.Atoi(x)

		if err != nil {
			continue
		}

		res = append(res, xInt)
	}

	return res[0], res[len(res)-1]
}
