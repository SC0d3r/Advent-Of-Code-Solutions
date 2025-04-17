package main

// What i have learned here is amazing
// 1. fast digit count O(1)
// (in js) Math.floor(Math.log10(x)) + 1
// 2. fast num split O(1)
// left part = num / Math.pow(10, howManyToSplit)
// right part = num % Math.pow(10, howManyToSplit)
// 3. cause we only need count we can use histogram (map)
// and for each one key (here a new stone) add a key (stone ) to that map
// and then sum how many times that we have seen that result
// the trick here for example if we have stone 1,1 then in that map we have
// {1: 2} then when we compute the next we will have {2024: 2}
// cause we do next[1 * 2024] += cnt which in here cnt for 1 is 2
// this keeps the memory footprint small

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const blinks = 75

func main() {
	// 1) read input into a map[value]count
	f, err := os.Open("2024/day11/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	counts := make(map[uint64]uint64)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		v, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		counts[v]++
	}

	// precompute powers of ten up to, say, 40 digits
	pow10 := make([]uint64, 41)
	pow10[0] = 1
	for i := 1; i < len(pow10); i++ {
		pow10[i] = pow10[i-1] * 10
	}

	// 2) blink 75 times, but only update the counts map
	for i := 0; i < blinks; i++ {
		next := make(map[uint64]uint64, len(counts))
		for v, cnt := range counts {
			switch {
			case v == 0:
				next[1] += cnt

			case digitCount(v)%2 == 0:
				d := digitCount(v) / 2
				div := pow10[d]
				next[v/div] += cnt
				next[v%div] += cnt

			default:
				next[v*2024] += cnt
			}
		}
		counts = next
	}

	// 3) sum up how many stones
	var total uint64
	for _, cnt := range counts {
		total += cnt
	}
	fmt.Printf("After %d blinks you have %d stones\n", blinks, total)
}

// digitCount returns ⌊log₁₀(v)⌋+1 (works for v>0)
func digitCount(v uint64) int {
	// using math.Log10 is slightly slower but very concise:
	return int(math.Floor(math.Log10(float64(v)))) + 1
}
