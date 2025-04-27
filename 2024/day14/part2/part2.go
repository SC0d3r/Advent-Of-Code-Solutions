package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Vel struct {
	x int
	y int
}

type Bot struct {
	p Pos
	v Vel
}

func main() {
	file, err := os.Open("2024/day14/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	bots := []*Bot{}

	lines := strings.SplitSeq(string(data), "\n")
	for l := range lines {
		b := createBot(l)
		bots = append(bots, &b)
	}

	log.Println("len", len(bots))

	// rows := 7
	// cols := 11

	rows := 103
	cols := 101

	// fb := Bot{p: Pos{x: 2, y: 4}, v: Vel{x: 2, y: -3}}
	// log.Println("fb before", fb)
	// step(&fb, rows, cols)
	// step(&fb, rows, cols)
	// step(&fb, rows, cols)
	// step(&fb, rows, cols)
	// step(&fb, rows, cols)
	// log.Println("fb after", fb)

	// for i := 0; i < 100; i++ {
	// 	for _, b := range bots {
	// 		step(b, rows, cols)
	// 	}
	// }

	// log.Println("bots after step", bots)

	// b := Bot{p: Pos{x: cols - 1, y: rows - 1}, v: Vel{}}
	// log.Println("inQuad", "bot", b, inQuad(b, rows, cols))
	tree := tposes(len(bots), rows, cols)
	// log.Println("tree", tree)

	for !isTree(bots, tree) {
		for _, b := range bots {
			// log.Println("steping")
			step(b, rows, cols)
		}
	}

	mx := cols / 2
	my := rows / 2
	log.Println("mx", mx, "my", my)
}

func step(b *Bot, rows int, cols int) {
	b.p.x = calc(b.p.x, b.v.x, cols)
	b.p.y = calc(b.p.y, b.v.y, rows)
}

func calc(p int, v int, t int) int {
	r := p + v
	return ((r % t) + t) % t
}

func createBot(l string) Bot {
	r := strings.Split(strings.TrimSpace(l), " ")
	p, v := r[0], r[1]

	ps := strings.Split(strings.Split(p, "=")[1], ",")
	p1, _ := strconv.Atoi(ps[0])
	p2, _ := strconv.Atoi(ps[1])
	pos := Pos{x: p1, y: p2}

	vs := strings.Split(strings.Split(v, "=")[1], ",")
	v1, _ := strconv.Atoi(vs[0])
	v2, _ := strconv.Atoi(vs[1])
	vel := Vel{x: v1, y: v2}

	return Bot{p: pos, v: vel}
}

func fibo(n int) int {
	l := []int{1, 1}
	for ; n > 0; n-- {
		l[0], l[1] = l[1], l[0]+l[1]
	}
	return l[1]
}

func tposes(lenBots int, rows int, cols int) map[int][]Pos {
	mx := cols / 2
	// my := rows / 2

	res := map[int][]Pos{}
	res[0] = []Pos{Pos{mx, 0}}

	for i := 0; i < rows; i++ {
		if lenBots == 0 {
			// no more bots left
			break
		}

		// pposes := res[len(res)-1]
		pposes := res[i]

		// gen next poses
		nposes := []Pos{}
		l := len(pposes) / 2
		// for i, pp := range pposes {
		pp := pposes[l] // the middle one
		for j := -(pp.y + 1); j <= pp.y+1; j++ {
			nposes = append(nposes, Pos{x: pp.x + j, y: pp.y + 1})
			lenBots--
		}

		res[i+1] = nposes
	}

	return res
}

func isTree(bots []*Bot, tree map[int][]Pos) bool {
	//     *
	//   * * *
	// * * * * *

	// so maybe
	// create all the poses for valid tree [][]Pos like {{{3,0}}, {{3,5},{8,5}}}
	// create map of bots based on theirs pos.y
	// for each row of valid poses:
	// 1. threre should be enough bots
	// 2. bots should be in specified cols (x)

	// used bots
	used := 0
	usedBs := []Bot{}

	// check each row of tree
	for /*row*/ _, trow := range tree {
		rowSatisfied := len(trow)
		for _, pos := range trow {
			// find a robot which satisfies this
			// if no bot found then return false
			fnd := false
			for _, b := range bots {
				if b.p.x == pos.x && b.p.y == pos.y {
					fnd = true
					used++
					rowSatisfied--
					usedBs = append(usedBs, *b)
					// should i break?
					// cause of overlaping bots
					// i have to check and use all bots right?
					break
				}
			}
			if !fnd {
				return false
			}
			// log.Println("trow satisfied", trow, "bots", usedBs)
			// log.Println("rowSatisfied", rowSatisfied)
		}

		if rowSatisfied > 0 {
			log.Println("didnd row satisfied")
			return false
		}
	}

	remBots := len(bots) - used
	log.Println("rembots", remBots, "used", used)
	if remBots > 0 {
		// still bots which are remaining
	}

	return true
}
