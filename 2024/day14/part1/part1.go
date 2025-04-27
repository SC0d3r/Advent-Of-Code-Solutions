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

	log.Println("bots", bots)

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

	for i := 0; i < 100; i++ {
		for _, b := range bots {
			step(b, rows, cols)
		}
	}

	log.Println("bots after step", bots)

	// b := Bot{p: Pos{x: cols - 1, y: rows - 1}, v: Vel{}}
	// log.Println("inQuad", "bot", b, inQuad(b, rows, cols))

	// is in top left quadrant?
	tlq := []Bot{}
	trq := []Bot{}

	blq := []Bot{}
	brq := []Bot{}

	mx := cols / 2
	my := rows / 2
	log.Println("mx", mx, "my", my)
	for _, b := range bots {
		// if inQuad(b, rows, cols) {
		// 	res = append(res, b)
		// }
		if b.p.x == mx || b.p.y == my {
			log.Println("on quad", b)
			continue
		}

		// top left quad?
		left := b.p.x < mx
		top := b.p.y < my

		if left && top {
			tlq = append(tlq, *b)
		} else if left && !top {
			blq = append(blq, *b)
		} else if !left && top {
			trq = append(trq, *b)
		} else if !left && !top {
			brq = append(brq, *b)
		} else {
			log.Println("bot is in the quad line", b)
		}
	}

	log.Println("tlq", tlq)
	log.Println("trq", trq)
	log.Println("blq", blq)
	log.Println("brq", brq)

	res := max(len(tlq), 1) * max(len(trq), 1) * max(1, len(blq)) * max(1, len(brq))
	log.Println("res", res)
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
