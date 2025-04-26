package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Partial struct {
	gcd int
	u   int
	v   int
}

type LB string

const (
	lbA LB = "A"
	lbB LB = "B"
)

type Btn struct {
	l    LB
	x, y int
}

type PZ struct {
	x, y int
}

type Press struct {
	btnA, btnB int
}

type MN struct {
	a, b Btn
	pz   PZ
}

func main() {
	file, err := os.Open("2024/day13/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var mns []MN
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		// each block is 4 lines: A, B, target, blank
		mns = append(mns, createMN(sc))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	sumTokens := 0
	for _, mn := range mns {
		pressA := newAlgo(mn.a, mn.b, mn.pz)
		pressB := newAlgo(mn.b, mn.a, mn.pz)

		if pressA.btnA < 0 || pressB.btnA < 0 {
			continue
		}

		ta := pressA.btnA*3 + pressA.btnB*1
		tb := pressB.btnA*1 + pressB.btnB*3
		best := ta
		if pressA.btnA != pressB.btnA && pressA.btnB != pressB.btnB {
			if tb < ta {
				best = tb
			}
		}
		sumTokens += best
	}

	log.Println("sum tokens:", sumTokens)
}

func createMN(sc *bufio.Scanner) MN {
	// line A
	a := strings.TrimSpace(sc.Text())
	btnA := readBtn(a, lbA)
	// line B
	sc.Scan()
	btnB := readBtn(strings.TrimSpace(sc.Text()), lbB)
	// line target
	sc.Scan()
	pz := readPz(strings.TrimSpace(sc.Text()))
	// blank (or EOF)
	sc.Scan()
	return MN{a: btnA, b: btnB, pz: pz}
}

func readBtn(line string, label LB) Btn {
	re := regexp.MustCompile(`\+(\d+)`)
	fields := re.FindAllStringSubmatch(line, -1)
	x, _ := strconv.Atoi(fields[0][1])
	y, _ := strconv.Atoi(fields[1][1])
	return Btn{l: label, x: x, y: y}
}

func readPz(line string) PZ {
	re := regexp.MustCompile(`=(\d+)`)
	fields := re.FindAllStringSubmatch(line, -1)
	x, _ := strconv.Atoi(fields[0][1])
	y, _ := strconv.Atoi(fields[1][1])
	return PZ{x: x + 10000000000000, y: y + 10000000000000}
}

// Extended Euclidean: returns gcd and one pair (u,v) such that u*x + v*y = gcd.
func gcd(x, y int) Partial {
	u0, u1 := 1, 0
	v0, v1 := 0, 1
	for y != 0 {
		q := x / y
		x, y = y, x-q*y
		u0, u1 = u1, u0-q*u1
		v0, v1 = v1, v0-q*v1
	}
	return Partial{gcd: x, u: u0, v: v0}
}

// Solve simultaneously:
//
//	m*a.x + n*b.x = pz.x
//	m*a.y + n*b.y = pz.y
//
// Return (-1,-1) if no non-negative integer solution.
func newAlgo(a, b Btn, pz PZ) Press {
	// 1) Check solvability of each axis
	gx := gcd(a.x, b.x)
	if pz.x%gx.gcd != 0 {
		return Press{-1, -1}
	}
	gy := gcd(a.y, b.y)
	if pz.y%gy.gcd != 0 {
		return Press{-1, -1}
	}

	// 2) One particular solution for the X-equation:
	//    m0 = u*(pz.x/g)
	//    n0 = v*(pz.x/g)
	m0 := gx.u * (pz.x / gx.gcd)
	n0 := gx.v * (pz.x / gx.gcd)

	// 3) We know the general X-soln is
	//    m = m0 + k*(b.x/gx)
	//    n = n0 - k*(a.x/gx)
	// Plug into Y:
	//   (m0 + k*(b.x/gx))*a.y + (n0 - k*(a.x/gx))*b.y = pz.y
	// ⇒ k * [ (b.x/gx)*a.y - (a.x/gx)*b.y ] = pz.y - (m0*a.y + n0*b.y)
	det := (b.x/gx.gcd)*a.y - (a.x/gx.gcd)*b.y
	rhs := pz.y - (m0*a.y + n0*b.y)

	switch {
	case det == 0:
		// either no solution (rhs≠0) or infinite (rhs=0).
		if rhs != 0 {
			return Press{-1, -1}
		}
		// infinite many—pick the one that is non-negative if m0,n0 already ≥0
		if m0 >= 0 && n0 >= 0 {
			return Press{m0, n0}
		}
		return Press{-1, -1}
	case rhs%det != 0:
		return Press{-1, -1}
	}

	k := rhs / det
	m := m0 + k*(b.x/gx.gcd)
	n := n0 - k*(a.x/gx.gcd)
	if m < 0 || n < 0 {
		return Press{-1, -1}
	}
	return Press{m, n}
}
