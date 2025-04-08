package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./2024/day3/inp1.txt")

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	res := parse(string(data))

	sum := 0
	for _, r := range res {
		sum += r
	}

	log.Println("sum is", sum)
}

func parse(data string) []int {
	cur := 0

	res := make([]int, 0)

	i := 0
	disabled := false
	for {
		// safty guard
		if i > len(data) {
			log.Println("this should not happen")
			break
		}
		i++

		// log.Println("cur is", cur)
		if cur >= len(data) {
			break
		}

		c, num, dis, err := step(cur, disabled, data)
		// log.Println("dis is", dis)
		disabled = dis
		cur = c
		if err != nil {
			continue
		}

		if !disabled {
			res = append(res, num)
		}
	}

	return res
}

func step(cur int, disabled bool, data string) (int, int, bool, error) {
	if cur >= len(data) {
		log.Fatal("out of range")
	}

	curCmd, found, shouldDisable := stepCMD(cur, data)

	cur = stepMul(cur, data)

	// log.Println("curCmd", curCmd, "cur", cur, "found", found, "shouldDis", shouldDisable)

	if found && curCmd < cur {
		// here the command is before mul so
		// we have to apply it
		disabled = shouldDisable
	}

	if disabled {
		return cur, 0, disabled, nil
	}

	// if it couldnt find a mul and we reach the end
	if cur >= len(data) {
		return cur, 0, disabled, errors.New("reached end")
	}

	cur, firstNum, err := stepNum(cur, data)
	if err != nil {
		// this mul is invalid
		return cur, 0, disabled, errors.New("invalid mul")
	}
	log.Println("first num is", firstNum)

	cur, secondNum, err := stepNum(cur, data)
	if err != nil {
		// this mul is invalid
		return cur, 0, disabled, errors.New("invalid mul")
	}
	log.Println("second num is", secondNum)

	return cur, firstNum * secondNum, disabled, nil
}

func stepMul(cur int, data string) int {
	s := ""

	for i, c := range data[cur:] {
		if len(s) == 0 && c == 'm' {
			s += string(c)
			continue
		}

		n, err := next(s)
		if err != nil {
			// invalid next
			s = ""
			continue
		}

		if string(c) == n {
			s += string(c)

			if s == "mul(" {
				// found a valid mul
				return cur + i + 1 // + 1 for (
			}
		} else {
			// this mul is not valid
			s = ""
		}
	}

	// search the whole data and didnt find
	return len(data) // out of range
}

func nextCMDPartValid(s string, nextPart string) (bool, error) {
	if s == "" {
		return nextPart == "d", nil
	}

	if s == "d" {
		return nextPart == "o", nil
	}

	if s == "do" {
		return nextPart == "n" || nextPart == "(", nil
	}

	if s == "do(" {
		return nextPart == ")", nil
	}

	if s == "don" {
		return nextPart == "'", nil
	}

	if s == "don'" {
		return nextPart == "t", nil
	}

	if s == "don't" {
		return nextPart == "(", nil
	}

	if s == "don't(" {
		return nextPart == ")", nil
	}

	return false, errors.New("invalid next")
}

func next(s string) (string, error) {
	if s == "" {
		return "m", nil
	}

	if s == "m" {
		return "u", nil
	}

	if s == "mu" {
		return "l", nil
	}

	if s == "mul" {
		return "(", nil
	}

	return "", errors.New("invalid next")
}

func stepNum(cur int, data string) (int, int, error) {
	num := ""
	y := 0
	for i, v := range data[cur:] {
		log.Println("in step num", data[cur:], "cur", cur)
		y += i
		if v == ',' || v == ')' {
			if len(num) == 0 {
				return cur + i + 1, 0, errors.New("invalid number")
			}

			resNum, err := strconv.Atoi(num)
			if err != nil {
				return cur + i, 0, errors.New("invalid number")
			}
			return cur + i + 1, resNum, nil
		}

		_, err := strconv.Atoi(string(v))
		if err != nil {
			return cur + i, 0, errors.New("invalid number")
		}

		num += string(v)
		log.Println("in stepNum num is", num)
		if len(num) > 3 {
			return cur + i, 0, errors.New("invalid number")
		}
	}

	return cur + y, 0, errors.New("invalid number")
}

func stepCMD(cur int, data string) (int, bool, bool) {
	s := ""
	for i, v := range data[cur:] {
		if s == "" {
			if v == 'd' {
				s += "d"
			}
			continue
		}

		// here s has atleast 'd'
		isValid, err := nextCMDPartValid(s, string(v))
		if err != nil {
			// invalid next
			s = ""
			continue
		}

		if isValid {
			s += string(v)

			if s == "do()" {
				// found a valid cmd
				return cur + i, true, false // + 1 for (
			}

			if s == "don't()" {
				// found a valid cmd
				return cur + i, true, true // + 1 for (
			}
		} else {
			// this cmd is not valid
			s = ""
		}
	}

	//         found?, shouldDisable?
	return cur, false, false
}
