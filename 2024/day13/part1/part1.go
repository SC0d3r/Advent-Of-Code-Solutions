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

	fmn := mns[3]
	press := algo(fmn)
	log.Println("press", press, "machine", fmn)

	sumt := 0
	for _, mn := range mns {
		press := algo(mn)
		if press.btnA == -1 {
			continue
		}
		sumt += (press.btnA * 3) + (press.btnB * 1)
	}

	log.Println("sum token", sumt)
}

func algo(mn MN) Press {
	// create a map of all possible btn a presses
	aps := []BtnPress{}        // [[val, count press]]
	bmap := map[int]BtnPress{} // map of value to how many presses

	// log.Println("pz is", mn.pz)

	if mn.a.x >= mn.b.x {

		// btn A
		vx := mn.a.x
		// vy := mn.a.y
		cnt := 1
		for vx < mn.pz.x {
			// amap[vx] = cnt
			bp := BtnPress{btn: mn.a, cnt: cnt}
			aps = append(aps, bp)
			cnt++
			vx += mn.a.x
			// vy = vy * cnt
		}
		// log.Println("vx is", vx, "cnt is", cnt)

		// btn B
		vx = mn.b.x
		// vy := mn.b.y
		cnt = 1
		for vx < mn.pz.x {
			bmap[vx] = BtnPress{btn: mn.b, cnt: cnt}
			cnt++
			vx += mn.b.x
			// vy = vy * cnt
		}

	} else {
		// btn B
		vx := mn.b.x
		// vy := mn.a.y
		cnt := 1
		for vx < mn.pz.x {
			// amap[vx] = cnt
			bp := BtnPress{btn: mn.b, cnt: cnt}
			aps = append(aps, bp)
			cnt++
			vx += mn.b.x
			// vy = vy * cnt
		}

		// btn A
		vx = mn.a.x
		// vy := mn.b.y
		cnt = 1
		for vx < mn.pz.x {
			bmap[vx] = BtnPress{btn: mn.a, cnt: cnt}
			cnt++
			vx += mn.a.x
			// vy = vy * cnt
		}
	}

	// log.Println("aps", aps)
	// log.Println("bmap", bmap)

	for i := len(aps) - 1; i >= 0; i-- {
		bp := aps[i]
		x := bp.btn.x * bp.cnt
		y := bp.btn.y * bp.cnt

		if x == mn.pz.x && y == mn.pz.y {
			// x,y on prize (found it)
			if mn.a.x > mn.b.x {
				return Press{
					btnA: bp.cnt,
					btnB: 0,
				}
			} else {
				return Press{
					btnB: bp.cnt,
					btnA: 0,
				}
			}
		}

		rem := mn.pz.x - x
		obp, exists := bmap[rem]
		if exists {
			// found a match
			// does y match too?
			resY := y + (obp.btn.y * obp.cnt)
			if resY == mn.pz.y {
				// found match
				if mn.a.x > mn.b.x {
					return Press{btnA: bp.cnt, btnB: obp.cnt}
				} else {
					return Press{btnB: bp.cnt, btnA: obp.cnt}
				}
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
	// log.Println("l", l, "btna", btna)

	// advance the scanner
	if !sc.Scan() {
		log.Fatal("shouldnt happen")
	}

	l = strings.TrimSpace(sc.Text())
	if l == "" {
		log.Fatal("this should not happen")
	}
	btnb := readBtn(l, lbB)
	// log.Println("l", l, "btnb", btnb)

	// advance the scanner
	if !sc.Scan() {
		log.Fatal("shouldnt happen")
	}

	l = strings.TrimSpace(sc.Text())
	if l == "" {
		log.Fatal("this should not happen")
	}
	pz := readPz(l)
	// log.Println("l", l, "pz", pz)

	// advance the scanner
	sc.Scan()

	return MN{a: btna, b: btnb, pz: pz}
}

func readPz(l string) PZ {
	// re := regexp.MustCompile(`((?<=\+)\d+)`)
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
	// re := regexp.MustCompile(`((?<=\+)\d+)`)
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
