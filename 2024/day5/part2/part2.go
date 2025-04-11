package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./2024/day5/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// rules[x][y]==true means x should come before y
	rules := make(map[int]map[int]bool)
	updates := [][]int{}

	// Read rules until a blank line is encountered
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		addRule(line, rules)
	}

	// Read update lines after the blank line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		update, err := parseUpdate(line)
		if err != nil {
			log.Fatal(err)
		}
		updates = append(updates, update)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	// Identify wrong updates (updates that don't follow the rules)
	var wrongUpdates [][]int
	for _, update := range updates {
		if !isUpdateCorrect(update, rules) {
			wrongUpdates = append(wrongUpdates, update)
		}
	}

	// Fix wrong updates by swapping conflicting elements until no changes occur
	for _, update := range wrongUpdates {
		for {
			changed := false
			for i := 0; i < len(update)-1; i++ {
				if valid, pair := checkRule(update[i], update[i+1:], rules); !valid {
					swap(update, pair[0], pair[1])
					changed = true
				}
			}
			if !changed {
				break
			}
		}
	}

	// Compute the middle element from each wrong update and sum them up
	mids := make([]int, 0, len(wrongUpdates))
	sum := 0
	for _, upd := range wrongUpdates {
		mid := upd[len(upd)/2] // using integer division for the mid index
		mids = append(mids, mid)
		sum += mid
	}

	log.Println("wrong updates", wrongUpdates)
	log.Println("mids", mids, "sum", sum)
	log.Println("took", time.Since(start))
}

// addRule adds a rule from a line "X|Y", meaning X should come before Y.
// It also sets the inverse rule (Y|X) to false.
func addRule(line string, rules map[int]map[int]bool) {
	parts := strings.Split(strings.TrimSpace(line), "|")
	if len(parts) != 2 {
		log.Fatal("invalid rule:", line)
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	// Set rule: x should be before y.
	if rules[x] == nil {
		rules[x] = make(map[int]bool)
	}
	rules[x][y] = true

	// Set inverse rule: y should not come before x.
	if rules[y] == nil {
		rules[y] = make(map[int]bool)
	}
	if valid, exists := rules[y][x]; exists && valid {
		log.Fatalf("contradiction: rule %d|%d conflicts with %d|%d", y, x, x, y)
	}
	rules[y][x] = false
}

// parseUpdate converts a comma-separated line into a slice of integers.
func parseUpdate(line string) ([]int, error) {
	parts := strings.Split(line, ",")
	result := make([]int, len(parts))
	for i, p := range parts {
		p = strings.TrimSpace(p)
		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}

// isUpdateCorrect checks whether an update order follows all the rules.
func isUpdateCorrect(update []int, rules map[int]map[int]bool) bool {
	for i := 0; i < len(update)-1; i++ {
		if valid, _ := checkRule(update[i], update[i+1:], rules); !valid {
			return false
		}
	}
	return true
}

// checkRule verifies that for element X with following elements Ys the rule holds.
// It returns false and the pair [X, Y] if a rule violation is found.
func checkRule(X int, Ys []int, rules map[int]map[int]bool) (bool, []int) {
	if ruleMap, exists := rules[X]; exists {
		for _, Y := range Ys {
			if valid, ok := ruleMap[Y]; ok {
				if !valid {
					return false, []int{X, Y}
				}
			}
		}
	}
	return true, nil
}

// swap exchanges the positions of x and y in the update slice.
func swap(update []int, x, y int) {
	idxX, idxY := -1, -1
	for i, v := range update {
		if v == x && idxX == -1 {
			idxX = i
		}
		if v == y && idxY == -1 {
			idxY = i
		}
		if idxX != -1 && idxY != -1 {
			break
		}
	}
	if idxX != -1 && idxY != -1 {
		update[idxX], update[idxY] = update[idxY], update[idxX]
	}
}
