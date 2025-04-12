package main

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Key struct {
	x   int
	y   int
	dir string
}

const WALL = "#"
const PATH = "."

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
	sdir := UP
	for i, _ := range m {
		for j, y := range m[i] {
			if y == UP || y == RIGHT || y == DOWN || y == LEFT {
				up = []int{i, j}
				sdir = y
			}
		}
	}

	start := time.Now()

	// potential paths
	ppaths := getPPaths(m, up, sdir)

	// get the all the eligible squares for creating obstacles
	// els is a set of all eligible sqaures to create potential obstacle to create a loop
	els := make(map[Key]struct{})
	for _, key := range ppaths {
		if up[0] == key.x && up[1] == key.y {
			// first guard pos is not eligible
			continue
		}

		if m[key.x][key.y] == WALL {
			// wall is not eligible
			continue
		}

		// note that we have to put "" for the dir otherwise
		// we the visited {3,4,">"} is different than visited {3,4,"<"}
		// which cause we want to generate the potential path of the guard and not
		// the direction we have to remove the dir so two of them consider one
		// when we are making or eligible paths (which is a set)
		els[Key{key.x, key.y, ""}] = struct{}{}
	}

	loops := 0

	// collector
	foundLoop := make(chan bool)
	done := make(chan bool)
	go func() {
		for range foundLoop {
			loops++
			log.Println("found loop, loops are", loops)
		}

		done <- true
	}()

	var wg sync.WaitGroup

	log.Println("number of potential walls", len(els))

	for key, _ := range els {
		wg.Add(1)
		go func() {
			defer wg.Done()

			visited := make(map[Key]struct{})

			// the first pos
			visited[Key{up[0], up[1], sdir}] = struct{}{}

			// copying cause hasLoop changes them
			nmap := cloneMap(m)
			guardPos := []int{up[0], up[1]}

			// create an obstacle in key
			nmap[key.x][key.y] = WALL

			if hasLoop(nmap, visited, guardPos) {
				// loops++
				foundLoop <- true
			}

			// undo the obstacle
			// m[key.x][key.y] = PATH
		}()
	}

	wg.Wait()
	close(foundLoop)
	<-done

	// log.Println("visited squares", visited, "len is", len(visited))
	log.Println("loops", loops)
	log.Println("took", time.Since(start))
	// log.Println("els ", els)
}

func hasLoop(m [][]string, visited map[Key]struct{}, up []int) bool {
	for {
		sq, end := step(m, up)
		// log.Println("after step", "sq", sq, "end", end)

		if end {
			// here we are exiting the map so no loop
			return false
		}

		// check if this sq (key) already is visited before
		if _, exists := visited[sq]; exists {
			// found loop
			return true
		}

		visited[sq] = struct{}{}

		// swap
		// log.Println("m before swap", m)
		m[up[0]][up[1]], m[sq.x][sq.y] = m[sq.x][sq.y], m[up[0]][up[1]]
		// log.Println("m after swap", m)

		// update the user pos
		up[0], up[1] = sq.x, sq.y
	}
}

func step(m [][]string, up []int) (Key, bool) {
	if len(up) != 2 {
		log.Fatal("up should have 2 values")
	}

	npos, dir := nextPos(m, up)

	if npos[0] < 0 || npos[0] >= len(m) || npos[1] < 0 || npos[1] >= len(m[0]) {
		// going outside of the map
		return Key{npos[0], npos[1], dir}, true
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
		// log.Print("hitting wall")
		return step(m, up)
	}

	// log.Println("step", "up", up, "npos", npos, "dir", dir)

	return Key{npos[0], npos[1], dir}, false
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

	log.Println("m is", m, "up is", up)
	panic("invalid user: " + user)
}

func cloneMap(m [][]string) [][]string {
	res := make([][]string, len(m))
	for i, xs := range m {
		res[i] = []string{}
		for _, x := range xs {
			res[i] = append(res[i], x)
		}
	}

	return res
}

func getPPaths(m [][]string, up []int, sdir string) []Key {
	visited := make(map[Key]struct{})
	// Use the correct initial state including the guard direction.
	startKey := Key{up[0], up[1], sdir}
	visited[startKey] = struct{}{}

	nmap := cloneMap(m)
	// Create a separate guard state copy.
	guardPos := []int{up[0], up[1]}
	// Run simulation on a clean copy.
	hasLoop(nmap, visited, guardPos)

	keys := make([]Key, 0)
	for v := range visited {
		keys = append(keys, v)
	}
	return keys
}
