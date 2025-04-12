package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const AN = "#"

type Vec struct {
	x int
	y int
}

type Loc struct {
	origin Vec
	mh     Vec
}

func main() {
	file, err := os.Open("./2024/day8/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	m := [][]string{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		xs := strings.Split(strings.TrimSpace(line), "")
		m = append(m, xs)
	}

	log.Println("m is", m)

	start := time.Now()
	// algorithm:
	// 1. find all the anthenas which are at least 2 or more of same type
	anthenas := map[string][]Loc{}

	// this will populate the anthenas
	// MHs := map[Vec]struct{}{}
	MHs := getAllMHs(m, anthenas)

	log.Println("MHs", MHs)

	validAntinodesVecs := map[Vec]struct{}{}

	antinodes := 0
	rows := len(m)
	cols := len(m[0])
	for loc, _ := range MHs {
		// validate mh
		o := loc.origin
		mh := loc.mh
		i := o.x + mh.x
		j := o.y + mh.y

		if i < 0 || i >= rows || j < 0 || j >= cols {
			// invalid
			continue
		}
		// valid
		validAntinodesVecs[Vec{i, j}] = struct{}{}
		antinodes++
	}

	log.Println("valid antinodes", validAntinodesVecs)
	log.Println("total", antinodes, "unique total", len(validAntinodesVecs))
	log.Println("took", time.Since(start))

	// 2. apply the manhattan distances on origina and product antinode locations
	// antinodes := map[Vec]struct{}{}
	// populateAntinodes(anthenas, antinodes)
	// print(m, validAntinodesVecs)
}

func getAllMHs(m [][]string, anthenas map[string][]Loc) map[Loc]struct{} {
	// set
	mhs := map[Loc]struct{}{}

	for i, xs := range m {
		for j, x := range xs {
			if x == AN || x == "." {
				// "#" antinode in Map
				continue
			}

			// if _, ok := strconv.ParseFloat(x); ok {
			// found anthena
			anthenas[x] = append(anthenas[x], Loc{
				origin: Vec{x: i, y: j},
				mh:     Vec{0, 0},
			})
		}
	}

	// populate the manhattan distances
	log.Println("anthenas", anthenas)
	for _, locs := range anthenas {
		if len(locs) <= 1 {
			// just one anthena with that type
			// so it doesnt create any antinodes
			continue
		}

		// here we have to popluate all of the manhattan distances
		log.Println("locs", locs)

		for i, xloc := range locs {
			for _, yloc := range locs[i+1:] {
				mhvec := calcMH(xloc.origin, yloc.origin)

				mhs[Loc{mh: mhvec, origin: xloc.origin}] = struct{}{}
				mhs[Loc{mh: Vec{-mhvec.x, -mhvec.y}, origin: yloc.origin}] = struct{}{}
			}
		}
	}

	return mhs
}

func calcMH(xvec Vec, yvec Vec) Vec {
	return Vec{xvec.x - yvec.x, xvec.y - yvec.y}
}

// func updateMHs(anthenas[x], Vec{i, j})

func print(m [][]string, vecs map[Vec]struct{}) {
	for v, _ := range vecs {
		m[v.x][v.y] = AN
	}

	for _, xs := range m {
		fmt.Println(strings.Join(xs, " "))
	}
}
