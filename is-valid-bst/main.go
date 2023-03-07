package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func isBST(n *Node, min, max int) bool {
	if n == nil {
		return true
	}
	if n.value < min || n.value > max {
		return false
	}
	return isBST(n.left, min, n.value) && isBST(n.right, n.value+1, max)
}

func main() {
	bt := &Node{
		value: 5,
		left: &Node{
			value: 2,
			left: &Node{
				value: 1,
			},
			right: &Node{
				value: 3,
			},
		},
		right: &Node{
			value: 7,
			left: &Node{
				value: 6,
			},
			right: &Node{
				value: 8,
			},
		},
	}

	if !isBST(bt, math.MinInt64, math.MaxInt64) {
		fmt.Println("Is not a valid BST!")
		return
	}
	fmt.Println("Is a valid BST!")
}
