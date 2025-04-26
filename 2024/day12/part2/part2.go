package main

import (
	"io"
	"log"
	"os"
	"strings"
)

var UL = []int{-1, -1}
var UR = []int{-1, 1}
var DL = []int{1, -1}
var DR = []int{1, 1}
var UP = []int{-1, 0}
var DOWN = []int{1, 0}
var LEFT = []int{0, -1}
var RIGHT = []int{0, 1}

// Pos is your renamed Point, with an extra label field L.
type Pos struct {
	x, y int
	l    string
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
			sides := corners(nds, m)
			for _, n := range nds {
				m[n.x][n.y] = "-"
			}
			// log.Println(nds[0].l, "area", area, "sides", len(keep), keep)
			log.Println(nds[0].l, "area", area, "sides", sides)
			ret += area * sides
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

func vIdx(i int, j int, rows int, cols int) bool {
	invalid := i < 0 || j < 0 || i >= rows || j >= cols
	return !invalid
}

func corners(nds []Pos, m [][]string) int {

	res := 0
	for _, p := range nds {
		log.Println("corners", p)
		if diffUL(p, m) {
			log.Println("diffUL", p)
			res++
		}

		if diffDL(p, m) {
			log.Println("diffDL", p)
			res++
		}

		if diffUR(p, m) {
			log.Println("diffUR", p)
			res++
		}

		if diffDR(p, m) {
			log.Println("diffDR", p)
			res++
		}

	}

	return res
	// ispow2 := (res & (res - 1)) == 0
	// if ispow2 {
	// 	return res
	// }

	// // we make it to the closest pow of 2
	// return int(math.Pow(2, math.Ceil(float64(math.Log2(float64(res))))))
}

func diffDR(p Pos, m [][]string) bool {
	rows := len(m)
	cols := len(m[0])

	// top
	i := DOWN[0] + p.x
	j := DOWN[1] + p.y
	upSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// left
	i = RIGHT[0] + p.x
	j = RIGHT[1] + p.y
	leftSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	i = DR[0] + p.x
	j = DR[1] + p.y
	diagSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// convex corner if both ortho are absent,
	// concave corner if both ortho are present but diag is absent
	return (!upSame && !leftSame) || (upSame && leftSame && !diagSame)
}

func diffUR(p Pos, m [][]string) bool {
	rows := len(m)
	cols := len(m[0])

	// top
	i := UP[0] + p.x
	j := UP[1] + p.y
	upSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// left
	i = RIGHT[0] + p.x
	j = RIGHT[1] + p.y
	leftSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	i = UR[0] + p.x
	j = UR[1] + p.y
	diagSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// convex corner if both ortho are absent,
	// concave corner if both ortho are present but diag is absent
	return (!upSame && !leftSame) || (upSame && leftSame && !diagSame)
}

func diffDL(p Pos, m [][]string) bool {
	rows := len(m)
	cols := len(m[0])

	// top
	i := DOWN[0] + p.x
	j := DOWN[1] + p.y
	upSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// left
	i = LEFT[0] + p.x
	j = LEFT[1] + p.y
	leftSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	i = DL[0] + p.x
	j = DL[1] + p.y
	diagSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// convex corner if both ortho are absent,
	// concave corner if both ortho are present but diag is absent
	return (!upSame && !leftSame) || (upSame && leftSame && !diagSame)
}

func diffUL(p Pos, m [][]string) bool {
	rows := len(m)
	cols := len(m[0])

	// top
	i := UP[0] + p.x
	j := UP[1] + p.y
	upSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// left
	i = LEFT[0] + p.x
	j = LEFT[1] + p.y
	leftSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	i = UL[0] + p.x
	j = UL[1] + p.y
	diagSame := vIdx(i, j, rows, cols) && m[i][j] == p.l

	// convex corner if both ortho are absent,
	// concave corner if both ortho are present but diag is absent
	return (!upSame && !leftSame) || (upSame && leftSame && !diagSame)
}
