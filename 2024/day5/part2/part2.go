package main

import (
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
}

func main() {
	file, err := os.Open("./2024/day5/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make(map[string]bool) // "X|Y":true
	updates := make([][]int, 0)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			// here we are processing updates
			updates = getUpdates(lines[i+1:])
			break
		}
		// process rules
		addRule(line, rules)
	}

	log.Println("rules", rules)
	log.Println("updates", updates)

	// a := []int{61, 13, 29}
	// a := []int{75, 47, 61, 53, 29}
	// for i := 0; i < len(a)-1; i++ {
	// 	v := a[i]
	// 	if !applyRules(v, a[i+1:], rules) {
	// 		log.Println("wrong update")
	// 	}
	// }

	cUpdates := make([][]int, 0)

	for _, update := range updates {
		correct := true
		for i := 0; i < len(update)-1; i++ {
			v := update[i]
			if !applyRules(v, update[i+1:], rules) {
				log.Println("wrong update", update)
				correct = false
				break
			}
		}

		if correct {
			cUpdates = append(cUpdates, update)
		}
	}

	// get the mids
	mids := make([]int, 0)
	for _, cs := range cUpdates {
		idx := math.Floor(float64(len(cs) / 2))
		mids = append(mids, cs[int(idx)])
	}

	sum := 0
	for _, x := range mids {
		sum += x
	}

	log.Println("mids", mids, "sum", sum)
}

func applyRules(X int, Ys []int, rules map[string]bool) bool {
	for _, Y := range Ys {
		key := strconv.Itoa(X) + "|" + strconv.Itoa(Y)
		val, ok := rules[key]
		log.Println("key is", key, "val is", val, "ok is", ok)
		if !ok {
			continue
		}

		// here the rule exists
		if !val {
			// invalid order
			return false
		}
	}

	return true
}

func addRule(line string, rules map[string]bool) {
	xs := strings.Split(strings.TrimSpace(line), "|")

	// X|Y set to true
	key1 := xs[0] + "|" + xs[1]
	rules[key1] = true

	// Y|X set to false meaning that Y should not come before X
	// but what if it
	key2 := xs[1] + "|" + xs[0]

	val, ok := rules[key2]
	if ok && val {
		// if it exists and value is true then this is a contradiction
		// so both X wants to be before Y and Y wants to be before X
		// is this allowed?
		log.Fatal("contradiction", "key2", key2, "key1", key1)
	}

	rules[key2] = false
}

func getUpdates(lines []string) [][]int {
	res := make([][]int, 0)
	for _, line := range lines {
		xs := strings.Split(strings.TrimSpace(line), ",")
		ret := make([]int, 0)
		for _, v := range xs {
			// log.Println("v is", v, "xs is", xs, "len", len(xs))
			vi, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, vi)
		}
		res = append(res, ret)
	}

	return res
}
