package main

import (
	"bufio"
	"log"
	"math"
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

const lbA LB = "A"
const lbB LB = "B"

type Btn struct {
	l LB
	x int
	y int
}

type PZ struct {
	x, y int
}

type Press struct {
	btnA int
	btnB int
}

type BtnPress struct {
	btn Btn
	cnt int // count
}

type MN struct {
	a  Btn
	b  Btn
	pz PZ
}

func main() {
	file, err := os.Open("2024/day13/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mns := []MN{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		mns = append(mns, createMN(sc))
	}

	log.Println("mns", mns)

	sumt := 0

	// fmc := mns[0]
	// press := newAlgo(fmc.a, fmc.b, fmc.pz)
	// log.Println("press", press)

	for _, mn := range mns {
		pressA := newAlgo(mn.a, mn.b, mn.pz)
		// 	// reverse the btn
		pressB := algo(mn.b, mn.a, mn.pz)

		if pressA.btnA == -1 || pressB.btnA == -1 {
			continue
		}

		ta := (pressA.btnA * 3) + (pressA.btnB * 1)
		t := ta

		if pressA.btnA != pressB.btnA && pressA.btnB != pressB.btnB {
			// if different
			tb := (pressB.btnA * 1) + (pressB.btnB * 3)
			t = min(ta, tb)
		}

		log.Println("pressA", pressA, "pressB", pressB, "ta", ta, "t", t)

		sumt += t
	}

	log.Println("sum token", sumt)
}

func createMN(sc *bufio.Scanner) MN {
	l := strings.TrimSpace(sc.Text())
	if l == "" {
		log.Fatal("this should not happen")
	}
	btna := readBtn(l, lbA)

	// advance the scanner
	if !sc.Scan() {
		log.Fatal("shouldnt happen")
	}

	l = strings.TrimSpace(sc.Text())
	if l == "" {
		log.Fatal("this should not happen")
	}
	btnb := readBtn(l, lbB)

	// advance the scanner
	if !sc.Scan() {
		log.Fatal("shouldnt happen")
	}

	l = strings.TrimSpace(sc.Text())
	if l == "" {
		log.Fatal("this should not happen")
	}
	pz := readPz(l)

	// advance the scanner
	sc.Scan()

	return MN{a: btna, b: btnb, pz: pz}
}

func readPz(l string) PZ {
	re := regexp.MustCompile(`\=(\d+)`)

	// btn a
	res := re.FindAllString(l, -1)
	xstr := res[0][1:] // skipping +
	ystr := res[1][1:]
	x, err := strconv.Atoi(xstr)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		log.Fatal(err)
	}

	// return PZ{x: x + 10000000000000, y: y + 10000000000000}
	return PZ{x: x, y: y}
}

func readBtn(l string, label LB) Btn {
	re := regexp.MustCompile(`\+(\d+)`)

	// btn a
	res := re.FindAllString(l, -1)
	xstr := res[0][1:] // skipping +
	ystr := res[1][1:]
	x, err := strconv.Atoi(xstr)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		log.Fatal(err)
	}

	return Btn{l: label, x: x, y: y}
}

func gcd(x int, y int) Partial {
	r0 := x
	r1 := y
	u0 := 1
	u1 := 0
	v0 := 0
	v1 := 1

	for r1 != 0 {
		q := int(math.Floor(float64(r0 / r1)))
		r0, r1 = r1, r0-q*r1
		u0, u1 = u1, u0-q*u1
		v0, v1 = v1, v0-q*v1
	}

	return Partial{
		gcd: r0,
		u:   u0,
		v:   v0,
	}
}

func newAlgo(a Btn, b Btn, pz PZ) Press {
	// for x
	gx := gcd(a.x, b.x)
	if pz.x%gx.gcd != 0 {
		// no solution
		return Press{btnA: -1, btnB: -1}
	}

	// for y
	gy := gcd(a.y, b.y)
	if pz.y%gy.gcd != 0 {
		// no solution
		return Press{btnA: -1, btnB: -1}
	}

	// for x
	m0 := gx.u * (pz.x / gx.gcd)
	n0 := gx.v * (pz.x / gx.gcd)

	t := 0
	// find the first t in which mx and n are both positive
	mx := m0 + ((b.x / gx.gcd) * -t) // 1 = t
	nx := n0 - ((a.x / gx.gcd) * -t) // 1 = t

	prevMx := mx
	prevNx := nx
	sign := 1
	noPath := 1
	for mx < 0 || nx < 0 {
		if noPath >= 3 {
			log.Println(":: NO SOLUTION", "a", a, "b", b, "mx", mx, "nx", nx)
			return Press{-1, -1}
		}

		t += sign
		mx = m0 + ((b.x / gx.gcd) * -t)
		nx = n0 - ((a.x / gx.gcd) * -t)

		if mx < 0 && mx < prevMx {
			// becomming even more negative
			noPath++
			sign *= -1
		}

		if nx < 0 && nx < prevNx {
			// becomming even more negative
			noPath++
			sign *= -1
		}

		// if t > 1e6 {
		// 	return Press{-1, -1}
		// }
	}

	valx := mx*a.x + nx*b.x

	log.Println("a", a, "b", b, "pz", pz, "valx", valx, "mx", mx, "nx", nx)

	return Press{btnA: mx, btnB: nx}
}
