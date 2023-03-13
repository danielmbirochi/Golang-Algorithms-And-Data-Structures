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

func MaxConsecutiveBranch(n *Node) int {
	if n == nil {
		return 0
	}
	return Max(
		Consecutive(n.left, n.value, 1),
		Consecutive(n.right, n.value, 1),
	)
}

func Consecutive(n *Node, prev, length int) int {
	if n == nil {
		return length
	}
	if n.value == prev+1 {
		lenLeft := Consecutive(n.left, n.value, length+1)
		lenRight := Consecutive(n.right, n.value, length+1)
		return Max(lenLeft, lenRight)
	} else {
		lenLeft := Consecutive(n.left, n.value, 1)
		lenRight := Consecutive(n.right, n.value, 1)
		return Max(lenLeft, lenRight, length)
	}
}

func Max(vals ...int) int {
	max := math.MinInt
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	t := &Node{
		value: 0,
		left: &Node{
			value: 1,
			left: &Node{
				value: 1,
			},
			right: &Node{
				value: 2,
				left: &Node{
					value: 3,
				},
			},
		},
		right: &Node{
			value: 2,
			left: &Node{
				value: 1,
			},
			right: &Node{
				value: 3,
				right: &Node{
					value: 4,
				},
			},
		},
	}

	fmt.Printf("Longest consecutive branch length : %d\n", MaxConsecutiveBranch(t))
}
