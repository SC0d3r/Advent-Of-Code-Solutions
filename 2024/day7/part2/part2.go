package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	val    int
	left   *Node
	right  *Node
	center *Node
}

func main() {
	file, err := os.Open("./2024/day7/inp1.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	eqs := map[int][]int{}
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ":")
		total, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		eqs[total] = []int{}

		operands := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, op := range operands {
			x, err := strconv.Atoi(op)
			if err != nil {
				log.Fatal(err)
			}
			eqs[total] = append(eqs[total], x)
		}
	}

	start := time.Now()

	log.Println("eqs", eqs)

	ops := []func(int, int) int{cat, add, mul}
	// a := 190
	// tree := makeTree(eqs[a][0], eqs[a], ops)
	// log.Println("tree is", tree)
	// log.Println("matches", traverseTree(a, tree))

	matches := []int{}
	for res, operands := range eqs {
		_, leafs := makeTree(operands[0], operands, ops)
		// if res == 3267 {
		// 	log.Println("lafs", leafs)
		// }

		for _, v := range leafs {
			if res == v {
				log.Println("found a match", res, "operands", operands)
				matches = append(matches, res)
				break
			}
		}

		// if traverseTree(res, tree) {
		// 	log.Println("found a match", res, "operands", operands)
		// 	matches = append(matches, res)
		// }
		// time.Sleep(5 * time.Millisecond)
	}

	log.Println("all matches", matches)
	sum := 0
	for _, v := range matches {
		sum += v
	}
	log.Println("sum is", sum)
	log.Println("took", time.Since(start))
}

func add(a int, b int) int {
	return a + b
}
func mul(a int, b int) int {
	return a * b
}
func cat(a int, b int) int {
	res, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func makeTree(val int, operands []int, ops []func(int, int) int) (*Node, []int) {
	// Note: the operand order matters for cat operation
	leftVal := ops[0](val, operands[1])
	rightVal := ops[1](val, operands[1])
	centerVal := ops[2](val, operands[1])

	if len(operands) == 2 {
		root := Node{
			val: val,
			left: &Node{
				val:    leftVal,
				left:   nil,
				center: nil,
				right:  nil,
			},
			right: &Node{
				val:    rightVal,
				left:   nil,
				center: nil,
				right:  nil,
			},
			center: &Node{
				val:    centerVal,
				left:   nil,
				center: nil,
				right:  nil,
			},
		}

		return &root, []int{leftVal, rightVal, centerVal}
	}

	leftTree, lls := makeTree(leftVal, operands[1:], ops)
	rightTree, rls := makeTree(rightVal, operands[1:], ops)
	centerTree, cls := makeTree(centerVal, operands[1:], ops)

	root := Node{
		val: val,
		left: &Node{
			val:    leftVal,
			left:   leftTree,
			right:  leftTree,
			center: leftTree,
			// left:   makeTree(leftVal, operands[1:], ops),
			// right:  makeTree(leftVal, operands[1:], ops),
			// center: makeTree(leftVal, operands[1:], ops),
		},
		right: &Node{
			val:    rightVal,
			left:   rightTree,
			right:  rightTree,
			center: rightTree,
			// left:   makeTree(rightVal, operands[1:], ops),
			// right:  makeTree(rightVal, operands[1:], ops),
			// center: makeTree(rightVal, operands[1:], ops),
		},
		center: &Node{
			val:    rightVal,
			left:   centerTree,
			right:  centerTree,
			center: centerTree,
			// left:   makeTree(centerVal, operands[1:], ops),
			// right:  makeTree(centerVal, operands[1:], ops),
			// center: makeTree(centerVal, operands[1:], ops),
		},
	}

	allLeafValues := []int{}
	allLeafValues = append(allLeafValues, lls...)
	allLeafValues = append(allLeafValues, rls...)
	allLeafValues = append(allLeafValues, cls...)
	return &root, allLeafValues
}

func traverseTree(match int, root *Node) bool {
	if root.val > match {
		// cause we only do + * || then we can return false here
		return false
	}

	if root.left == nil || root.right == nil || root.center == nil {
		// this is a leaf node
		// end of search
		return root.val == match
	}

	// if root.val == match {
	// 	return true
	// }

	return traverseTree(match, root.left) || traverseTree(match, root.right) || traverseTree(match, root.center)
}
