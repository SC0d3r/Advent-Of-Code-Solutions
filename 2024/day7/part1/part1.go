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
	val   int
	left  *Node
	right *Node
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

	// collector
	matches := []int{}

	ops := []func(int, int) int{add, mul}
	// a := 190
	// tree := makeTree(eqs[a][0], eqs[a], ops)
	// log.Println("tree is", tree)
	// log.Println("matches", traverseTree(a, tree))

	for res, operands := range eqs {
		tree := makeTree(operands[0], operands, ops)
		if traverseTree(res, tree) {
			log.Println("found a match", res, "operands", operands)
			matches = append(matches, res)
		}
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

func makeTree(val int, operands []int, ops []func(int, int) int) *Node {
	leftVal := ops[0](operands[1], val)
	rightVal := ops[1](operands[1], val)

	if len(operands) == 2 {
		root := Node{
			val: val,
			left: &Node{
				val:   leftVal,
				left:  nil,
				right: nil,
			},
			right: &Node{
				val:   rightVal,
				left:  nil,
				right: nil,
			},
		}

		return &root
	}

	root := Node{
		val: val,
		left: &Node{
			val:   leftVal,
			left:  makeTree(leftVal, operands[1:], ops),
			right: makeTree(leftVal, operands[1:], ops),
		},
		right: &Node{
			val:   rightVal,
			left:  makeTree(rightVal, operands[1:], ops),
			right: makeTree(rightVal, operands[1:], ops),
		},
	}

	return &root
}

func traverseTree(match int, root *Node) bool {
	if root.left == nil || root.right == nil {
		// this is a leaf node
		// end of search
		return root.val == match
	}

	// if root.val == match {
	// 	return true
	// }

	return traverseTree(match, root.left) || traverseTree(match, root.right)
}
