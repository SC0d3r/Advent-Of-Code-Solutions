package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
	l string
}

func main() {
	file, err := os.Open("2024/day12/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	m := make([][]string, 0)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		xs := strings.Split(strings.TrimSpace(line), "")
		m = append(m, xs)
	}

	log.Println("m", m)

	ret := 0
	for i, xs := range m {
		for j, x := range xs {
			// log.Println("before", m)
			if x == "-" {
				continue
			}

			nds := nodes(Pos{l: x, x: i, y: j}, m)
			// log.Println("nds", nds, "m", m)
			area := len(nds)
			sumpms := 0
			for _, p := range nds {
				// m[p.x][p.y] = "-"
				x := perimeter(p, m)
				// log.Println("per of", p, "is", x)
				sumpms += x
			}
			for _, n := range nds {
				m[n.x][n.y] = "-"
			}
			// log.Println(nds[0].l, "area", area, "sumpms", sumpms)
			ret += area * sumpms
		}
	}

	log.Println("ret", ret)
}

func nodes(p Pos, m [][]string) []Pos {
	nds := make([]Pos, 0)

	st := []Pos{}
	st = append(st, p)

	rows := len(m)
	cols := len(m[0])
	sides := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	mem := make(map[Pos]struct{})
	for len(st) > 0 {
		// pop stack
		h := st[len(st)-1]
		st = st[:len(st)-1]
		m[h.x][h.y] = "-"
		if _, exist := mem[h]; !exist {
			nds = append(nds, h)
		}
		mem[h] = struct{}{}

		// add nghs to stack
		for _, s := range sides {
			i := s[0] + h.x
			j := s[1] + h.y
			if vIdx(i, j, rows, cols) && m[i][j] == h.l {
				// valid same label ngh
				st = append(st, Pos{x: i, y: j, l: h.l})
			}
		}
	}

	// restore
	for _, n := range nds {
		m[n.x][n.y] = n.l
	}

	return nds
}

func perimeter(p Pos, m [][]string) int {
	rows := len(m)
	cols := len(m[0])
	sides := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	res := 4
	for _, s := range sides {
		i := p.x + s[0]
		j := p.y + s[1]
		if !vIdx(i, j, rows, cols) {
			continue
		}

		// valid idx
		if m[i][j] == p.l {
			// log.Println("eq", p, "i", i, "j", j)
			res -= 1
		}
	}

	return res
}

func vIdx(i int, j int, rows int, cols int) bool {
	invalid := i < 0 || j < 0 || i >= rows || j >= cols
	return !invalid
}
