package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Combined struct {
	p1 Pos
	p2 Pos
}

type Dir struct {
	left   bool
	right  bool
	top    bool
	bottom bool
}

type Pos struct {
	x   int
	y   int
	val int
}

func main() {
	file, err := os.Open("2024/day10/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	theads := []Pos{}

	m := [][]int{}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		xs := strings.Split(strings.TrimSpace(line), "")
		ret := []int{}
		for j, x := range xs {
			// if x == "." {
			// 	continue
			// }

			n, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}

			if n == 0 {
				theads = append(theads, Pos{i, j, 0})
			}

			ret = append(ret, n)
		}
		m = append(m, ret)
	}

	// theads = append(theads, Pos{2, 4, 0})

	log.Println("m", m)
	log.Println("theads", theads)

	// maps theads to how many reachable 9
	res := map[Pos]int{}

	rows := len(m)
	cols := len(m[0])

	nines := map[Combined]struct{}{}

	for _, th := range theads {
		stack := [][]Pos{}

		// init the stack
		stack = append(stack, []Pos{th})
		log.Println("stack", stack)

		for len(stack) > 0 {
			// pop the stack
			// current path
			cp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			log.Println("cp", cp, "stack", stack)

			ph := cp[len(cp)-1]

			// all next dirs
			left := Pos{ph.x, ph.y - 1, 0}
			right := Pos{ph.x, ph.y + 1, 0}
			top := Pos{ph.x - 1, ph.y, 0}
			bot := Pos{ph.x + 1, ph.y, 0}

			// valid paths
			if validIdx(left, rows, cols) && nicp(left, cp) {
				left.val = m[left.x][left.y]
				log.Println("in left")
				newPath := make([]Pos, len(cp), len(cp)+1)
				copy(newPath, cp)
				newPath = append(newPath, left)
				if finishedPath(newPath, m) {
					// found a valid
					log.Println("finshed for th", th, "finshed path", newPath)
					if _, exists := nines[Combined{th, left}]; exists {
						log.Println("this nine already exists for this th", th, "nine", left)
					} else {
						nines[Combined{th, left}] = struct{}{}
						res[th]++
					}
				} else {
					// not yet finished
					if !notValidPath(newPath, m) {
						log.Println("in left adding", "newPath", newPath, "stack", stack)
						stack = append(stack, newPath)
					} else {
						log.Println("left is not valid")
					}
				}
			}

			if validIdx(right, rows, cols) && nicp(right, cp) {
				right.val = m[right.x][right.y]
				log.Println("in right")
				newPath := make([]Pos, len(cp), len(cp)+1)
				copy(newPath, cp)
				newPath = append(newPath, right)
				if finishedPath(newPath, m) {
					log.Println("finshed for th", th, "finshed path", newPath)
					if _, exists := nines[Combined{th, right}]; exists {
						log.Println("this nine already exists for this th", th, "nine", right)
					} else {
						nines[Combined{th, right}] = struct{}{}
						res[th]++
					}
				} else {
					if !notValidPath(newPath, m) {
						log.Println("in right adding", "newPath", newPath, "stack", stack)
						stack = append(stack, newPath)
					} else {
						log.Println("right is not valid")
					}
				}
			}

			if validIdx(top, rows, cols) && nicp(top, cp) {
				top.val = m[top.x][top.y]
				log.Println("in top")
				newPath := make([]Pos, len(cp), len(cp)+1)
				copy(newPath, cp)
				newPath = append(newPath, top)
				if finishedPath(newPath, m) {
					log.Println("finshed for th", th, "finshed path", newPath)
					if _, exists := nines[Combined{th, top}]; exists {
						log.Println("this nine already exists for this th", th, "nine", top)
					} else {
						nines[Combined{th, top}] = struct{}{}
						res[th]++
					}
				} else {
					if !notValidPath(newPath, m) {
						log.Println("in top adding", "newPath", newPath, "stack", stack)
						stack = append(stack, newPath)
					} else {
						log.Println("top is not valid")
					}
				}
			}

			if validIdx(bot, rows, cols) && nicp(bot, cp) {
				bot.val = m[bot.x][bot.y]
				log.Println("in bot")
				newPath := make([]Pos, len(cp), len(cp)+1)
				copy(newPath, cp)
				newPath = append(newPath, bot)
				if finishedPath(newPath, m) {
					log.Println("finshed for th", th, "finshed path", newPath)
					if _, exists := nines[Combined{th, bot}]; exists {
						log.Println("this nine already exists for this th", th, "nine", bot)
					} else {
						nines[Combined{th, bot}] = struct{}{}
						res[th]++
					}
				} else {
					if !notValidPath(newPath, m) {
						log.Println("in bot adding", "newPath", newPath, "stack", stack)
						stack = append(stack, newPath)
					} else {
						log.Println("bot is not valid")
					}
				}
			}
		}
	}

	log.Println("res", res)
	sum := 0
	for _, v := range res {
		sum += v
	}

	log.Println("sum", sum)
}

func notValidPath(path []Pos, m [][]int) bool {
	if len(path) == 0 || len(path) > 10 {
		return true
	}

	pCurrent := path[len(path)-1]
	xCurrent := m[pCurrent.x][pCurrent.y]

	if len(path) == 1 && xCurrent != 0 {
		return true
	}

	pPrev := path[len(path)-2]
	xPrev := m[pPrev.x][pPrev.y]

	return xCurrent != (xPrev + 1)
}

func finishedPath(path []Pos, m [][]int) bool {
	// if len(path) == 0 || len(path) > 10 {
	// 	return true
	// }

	// pCurrent := path[len(path)-1]
	// xCurrent := m[pCurrent.x][pCurrent.y]

	// if len(path) == 1 && xCurrent != 0 {
	// 	return true
	// }

	// pPrev := path[len(path)-2]
	// xPrev := m[pPrev.x][pPrev.y]

	// return xCurrent == (xPrev + 1)

	if len(path) < 10 {
		return false
	}

	c := 0
	for _, p := range path {
		x := m[p.x][p.y]
		if c-x != 0 {
			return false
		}
		c++
	}

	return true
}

func validIdx(pos Pos, rows int, cols int) bool {
	return !(pos.x < 0 || pos.x >= rows || pos.y < 0 || pos.y >= cols)
}

func nicp(pos Pos, cp []Pos) bool {
	for _, p := range cp {
		if pos.x == p.x && pos.y == p.y {
			// log.Println("pos", pos, "cp", cp, "resturn false")
			return false
		}
	}

	// not in current path
	return true
}

func getDir(cp []Pos) Dir {
	p1 := cp[len(cp)-1]
	p2 := cp[len(cp)-2]

	left := (p1.x-p2.x) == 0 && (p1.y-p2.y) < 0
	right := (p1.x-p2.x) == 0 && (p1.y-p2.y) > 0

	top := (p1.x-p2.x) < 0 && (p1.y-p2.y) == 0
	bottom := (p1.x-p2.x) > 0 && (p1.y-p2.y) == 0

	return Dir{
		left,
		right,
		bottom,
		top,
	}
}
