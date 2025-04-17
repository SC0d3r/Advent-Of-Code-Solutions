package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/day11/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	stones := []int{}
	strStones := strings.Split(string(data), " ")
	for _, s := range strStones {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatal(err)
		}

		stones = append(stones, n)
	}

	log.Println("stones", stones)

	blinks := 25

	// first blink
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	log.Println("after", blinks, "blink", stones)
	log.Println("len(stones)", len(stones))
}

func blink(stones []int) []int {
	newStrones := []int{}
	for _, s := range stones {
		if s == 0 {
			// log.Println("replace by 1", "s", s)
			newStrones = append(newStrones, 1)
			continue
		}

		ss := strconv.Itoa(s)
		if len(ss)%2 == 0 {
			// here we have to split
			left, right := ss[:len(ss)/2], ss[len(ss)/2:]

			ln, err := strconv.Atoi(left)
			if err != nil {
				log.Fatal(err)
			}

			rn, err := strconv.Atoi(right)
			if err != nil {
				log.Fatal(err)
			}

			// log.Println("halfing", "ss", ss, "left", left, "right", right, "ln", ln, "rn", rn)
			newStrones = append(newStrones, ln)
			newStrones = append(newStrones, rn)
			continue
		}

		// log.Println("*2024", "s", s)
		newStrones = append(newStrones, 2024*s)
	}
	return newStrones
}
