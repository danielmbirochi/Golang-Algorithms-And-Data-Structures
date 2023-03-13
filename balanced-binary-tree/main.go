package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func IsBalancedTree(n *Node) bool {
	return BalancedHeight(n) > -1
}

func BalancedHeight(n *Node) int {
	if n == nil {
		return 0
	}

	h1 := BalancedHeight(n.left)
	h2 := BalancedHeight(n.right)

	if h1 == -1 || h2 == -1 {
		return -1
	}
	if AbsDiff(h1, h2) > 1 {
		return -1
	}

	if h1 > h2 {
		return h1 + 1
	}
	return h2 + 1
}

func AbsDiff(x, y int) int {
	x = abs(x)
	y = abs(y)
	if x < y {
		return y - x
	}
	return x - y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	t1 := &Node{
		value: 1,
		left: &Node{
			value: 2,
			left: &Node{
				value: 4,
				left: &Node{
					value: 8,
				},
			},
			right: &Node{
				value: 5,
			},
		},
		right: &Node{
			value: 3,
			left: &Node{
				value: 6,
			},
			right: &Node{
				value: 7,
			},
		},
	}

	fmt.Printf("Tree1 is balanced: %t\n", IsBalancedTree(t1))

	t2 := &Node{
		value: 1,
		left: &Node{
			value: 2,
			left: &Node{
				value: 4,
				left: &Node{
					value: 8,
				},
			},
			right: &Node{
				value: 5,
			},
		},
		right: &Node{
			value: 3,
			left: &Node{
				value: 6,
				left: &Node{
					value: 9,
					left: &Node{
						value: 10,
					},
				},
			},
			right: &Node{
				value: 7,
			},
		},
	}

	fmt.Printf("Tree2 is balanced: %t\n", IsBalancedTree(t2))
}
