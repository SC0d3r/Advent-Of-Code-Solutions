package main

import (
	"io"
	"log"
	"os"
	"strings"
)

// Pos is your renamed Point, with an extra label field L.
type Pos struct {
	x, y int
	l    string
}

// Vertex is a corner of a unit square.
type Vertex struct{ X, Y int }

// Edge is an undirected unit‐length segment between two vertices.
type Edge struct{ A, B Vertex }

// canonicalEdge orders the endpoints so we can use it as a map key.
func canonicalEdge(v1, v2 Vertex) Edge {
	if v1.X < v2.X || (v1.X == v2.X && v1.Y < v2.Y) {
		return Edge{A: v1, B: v2}
	}
	return Edge{A: v2, B: v1}
}

// dir returns 0=E,1=N,2=W,3=S for the unit step v1→v2.
func dir(v1, v2 Vertex) int {
	switch {
	case v2.X > v1.X:
		return 0
	case v2.Y > v1.Y:
		return 1
	case v2.X < v1.X:
		return 2
	default:
		return 3
	}
}

// countSides computes the number of straight sides in the boundary loop.
// We take cells as []Pos (ignoring the L field).
func countSides(cells []Pos) int {
	// 1) build set of boundary edges
	edges := make(map[Edge]bool)
	addEdge := func(v1, v2 Vertex) {
		e := canonicalEdge(v1, v2)
		if edges[e] {
			delete(edges, e) // shared → interior
		} else {
			edges[e] = true
		}
	}

	for _, c := range cells {
		v00 := Vertex{c.x, c.y}
		v10 := Vertex{c.x + 1, c.y}
		v11 := Vertex{c.x + 1, c.y + 1}
		v01 := Vertex{c.x, c.y + 1}
		addEdge(v00, v10) // bottom
		addEdge(v10, v11) // right
		addEdge(v11, v01) // top
		addEdge(v01, v00) // left
	}

	// 2) build adjacency list for the boundary graph
	adj := make(map[Vertex][]Vertex)
	for e := range edges {
		adj[e.A] = append(adj[e.A], e.B)
		adj[e.B] = append(adj[e.B], e.A)
	}

	// 3) pick a canonical start vertex: smallest (X,Y)
	var start Vertex
	first := true
	for v := range adj {
		if first || v.X < start.X || (v.X == start.X && v.Y < start.Y) {
			start, first = v, false
		}
	}

	// 4) initialize the walk
	nbrs := adj[start]
	if len(nbrs) != 2 {
		return 0 // not a single loop
	}
	firstNext := nbrs[0]
	prevDir := dir(start, firstNext)
	sides := 1

	prev := start
	current := firstNext

	// 5) walk until we close back to (start → firstNext),
	//    *without* counting that final turn.
	for {
		// pick the next vertex (the one that isn’t prev)
		neis := adj[current]
		var next Vertex
		if neis[0] == prev {
			next = neis[1]
		} else {
			next = neis[0]
		}

		// if we're about to walk back over the very first edge, stop here.
		if current == start && next == firstNext {
			break
		}

		// count a new side whenever direction changes
		d := dir(current, next)
		if d != prevDir {
			sides++
			prevDir = d
		}

		// advance
		prev, current = current, next
	}

	return sides
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
			sides := countSides(nds)
			for _, n := range nds {
				m[n.x][n.y] = "-"
			}
			// log.Println(nds[0].l, "area", area, "sides", len(keep), keep)
			// log.Println(nds[0].l, "area", area, "sides", sides)
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
