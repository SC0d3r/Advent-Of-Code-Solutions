package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
	for _, mn := range mns {
		pressA := algo(mn.a, mn.b, mn.pz)
		// reverse the btn
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

		log.Println("pressA", pressA, "pressB", pressB, "ta", ta)

		sumt += t
	}

	log.Println("sum token", sumt)
}

func algo(ba Btn, bb Btn, pz PZ) Press {
	// create a map of all possible btn a presses
	aps := []BtnPress{}        // [[val, count press]]
	bmap := map[int]BtnPress{} // map of value to how many presses

	// btn A
	vx := ba.x
	cnt := 1
	for vx < pz.x {
		bp := BtnPress{btn: ba, cnt: cnt}
		aps = append(aps, bp)
		cnt++
		vx += ba.x
	}
	// log.Println("vx is", vx, "cnt is", cnt)

	// btn B
	vx = bb.x
	cnt = 1
	for vx < pz.x {
		bmap[vx] = BtnPress{btn: bb, cnt: cnt}
		cnt++
		vx += bb.x
	}

	for i := len(aps) - 1; i >= 0; i-- {
		bp := aps[i]
		x := bp.btn.x * bp.cnt
		y := bp.btn.y * bp.cnt

		if x == pz.x && y == pz.y {
			// x,y on prize (found it)
			return Press{
				btnA: bp.cnt,
				btnB: 0,
			}
		}

		rem := pz.x - x
		obp, exists := bmap[rem]
		if exists {
			// found a match
			// does y match too?
			resY := y + (obp.btn.y * obp.cnt)
			if resY == pz.y {
				// found match
				return Press{btnA: bp.cnt, btnB: obp.cnt}
			}
		}
	}

	// not found a match
	return Press{btnA: -1, btnB: -1}
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
