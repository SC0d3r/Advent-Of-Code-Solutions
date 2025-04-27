package main

import (
	"fmt"
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

	rows := 103
	cols := 101

	for t := range 200000 {
		for _, b := range bots {
			step(b, rows, cols)
		}

		// no overlaping
		ovlp := false
		m := map[Pos]struct{}{}
		for _, b := range bots {
			_, exist := m[b.p]
			if exist {
				ovlp = true
				break
			}

			m[b.p] = struct{}{}
		}
		if !ovlp {
			// no overlaping
			plotBots(bots, rows, cols)
			log.Fatal(":: Found:" + " t:" + strconv.Itoa(t+1))
		}
	}

	plotBots(bots, rows, cols)
}

func plotBots(bots []*Bot, rows, cols int) {
	// Create an empty grid
	grid := make([][]rune, rows)
	for y := range grid {
		grid[y] = make([]rune, cols)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	// Place bots
	for _, b := range bots {
		grid[b.p.y][b.p.x] = '#'
	}

	// Print grid
	for _, row := range grid {
		fmt.Println(string(row))
	}
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
