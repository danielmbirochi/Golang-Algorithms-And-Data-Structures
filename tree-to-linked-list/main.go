package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func TreeToLinkedList(n *Node) *Node {
	if n == nil {
		return n
	}

	leftList := TreeToLinkedList(n.left)
	rightList := TreeToLinkedList(n.right)

	n.left = n
	n.right = n

	n = Concatenate(leftList, n)
	n = Concatenate(n, rightList)
	return n
}

// Concatenate receives 2 circular doubly linked lists and return
// a pointer to the first linked list concatenated w/ the second list.
func Concatenate(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	aEnd := a.left
	bEnd := b.left

	a.left = bEnd
	bEnd.right = a

	aEnd.right = b
	b.left = aEnd

	return a
}

// Traverse receives a circular doubly linked list
// and print all it nodes.
func Traverse(n *Node) {
	end := n.left

	currentNode := n
	for currentNode.value != end.value {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.right
	}
	fmt.Printf("%d ", currentNode.value)
}

func main() {
	t := &Node{
		value: 1,
		left: &Node{
			value: 2,
			left: &Node{
				value: 4,
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

	l := TreeToLinkedList(t)
	Traverse(l)
}
