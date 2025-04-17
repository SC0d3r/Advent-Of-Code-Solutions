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

	stones := []uint64{}
	strStones := strings.Split(string(data), " ")
	for _, s := range strStones {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatal(err)
		}

		stones = append(stones, uint64(n))
	}

	log.Println("stones", stones)

	blinks := 75

	buf := make([]uint64, 0)
	output := make([]uint64, 0)

	// first blink
	for i := 0; i < blinks; i++ {
		// clear content but reuse memory
		output = output[:0]

		buf = blink(stones, buf)

		output = append(output, buf...)
		stones, output = output, stones
	}

	log.Println("after", blinks, "blink", stones)
	log.Println("len(stones)", len(stones))
}

func blink(stones []uint64, buf []uint64) []uint64 {
	// clear it without realloc
	// reset slice
	buf = buf[:0]

	// newStrones := []uint64{}
	for _, s := range stones {
		if s == 0 {
			// log.Println("replace by 1", "s", s)
			buf = append(buf, 1)
			continue
		}

		// ss := strconv.Itoa(int(s))
		ss := strconv.FormatUint(s, 10)
		if len(ss)%2 == 0 {
			// here we have to split
			left, right := ss[:len(ss)/2], ss[len(ss)/2:]

			ln, err := strconv.ParseUint(left, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			rn, err := strconv.ParseUint(right, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// log.Println("halfing", "ss", ss, "left", left, "right", right, "ln", ln, "rn", rn)
			buf = append(buf, uint64(ln))
			buf = append(buf, uint64(rn))
			continue
		}

		// log.Println("*2024", "s", s)
		buf = append(buf, uint64(2024*s))
	}
	return buf
}
