package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const WALL = "#"

const (
	UP    = "^"
	RIGHT = ">"
	DOWN  = "v"
	LEFT  = "<"
)

func main() {
	file, err := os.Open("./2024/day6/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	m := [][]string{}

	for _, line := range lines {
		xs := strings.Split(strings.TrimSpace(line), "")
		m = append(m, xs)
	}

	// user post
	up := make([]int, 2, 2)
	for i, _ := range m {
		for j, y := range m[i] {
			if y == UP || y == RIGHT || y == DOWN || y == LEFT {
				up = []int{i, j}
			}
		}
	}

	type Key struct {
		x int
		y int
	}

	visited := make(map[Key]struct{})

	// the first pos
	visited[Key{up[0], up[1]}] = struct{}{}

	start := time.Now()

	// now we have the map
	for {
		sq, end := step(m, up)
		log.Println("after step", "sq", sq, "end", end)

		if end {
			break
		}

		visited[Key{sq[0], sq[1]}] = struct{}{}

		// swap
		// log.Println("m before swap", m)
		m[up[0]][up[1]], m[sq[0]][sq[1]] = m[sq[0]][sq[1]], m[up[0]][up[1]]
		// log.Println("m after swap", m)

		// update the user pos
		up = sq
	}

	log.Println("visited squares", visited, "len is", len(visited))
	log.Println("took", time.Since(start))
}

func step(m [][]string, up []int) ([]int, bool) {
	if len(up) != 2 {
		log.Fatal("up should have 2 values")
	}

	npos, dir := nextPos(m, up)

	if npos[0] < 0 || npos[0] >= len(m) || npos[1] < 0 || npos[1] >= len(m[0]) {
		// going outside of the map
		return npos, true
	}

	// we should check the rules
	// 1. if WALL then we have to turn 90 degrees
	sq := m[npos[0]][npos[1]]
	if sq == WALL {
		// do 90 degrees turn
		ndir := dir
		if dir == UP {
			ndir = RIGHT
		}
		if dir == RIGHT {
			ndir = DOWN
		}
		if dir == DOWN {
			ndir = LEFT
		}
		if dir == LEFT {
			ndir = UP
		}

		// change the player in map (m)
		m[up[0]][up[1]] = ndir
		log.Print("hitting wall")
		return step(m, up)
	}

	log.Println("step", "up", up, "npos", npos, "dir", dir)

	return npos, false
}

func nextPos(m [][]string, up []int) ([]int, string) {
	// going up
	user := m[up[0]][up[1]]
	if user == "" {
		log.Fatal("Couldnt find the user", m, up)
	}

	// going up
	if user == UP {
		return []int{up[0] - 1, up[1]}, UP
	}

	// going down
	if user == DOWN {
		return []int{up[0] + 1, up[1]}, DOWN
	}

	// going left
	if user == LEFT {
		return []int{up[0], up[1] - 1}, LEFT
	}

	// going down
	if user == RIGHT {
		return []int{up[0], up[1] + 1}, RIGHT
	}

	panic("invalid user: " + user)
}
