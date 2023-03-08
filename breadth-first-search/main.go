package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Queue struct {
	items []Node
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Enqueue(n Node) {
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() Node {
	if q.IsEmpty() {
		panic("queue underflow")
	}
	x := q.items[0]
	q.items = q.items[1:]
	return x
}

// Traverse func traverses the given binary tree in
// a breadth first search (BFS) approach.
func Traverse(n *Node) {
	if n == nil {
		fmt.Println("Tree has no element.")
		return
	}

	var q Queue
	q.Enqueue(*n)

	for !q.IsEmpty() {
		currentNode := q.Dequeue()
		fmt.Printf("%d ", currentNode.value)

		if currentNode.left != nil {
			q.Enqueue(*currentNode.left)
		}
		if currentNode.right != nil {
			q.Enqueue(*currentNode.right)
		}
	}
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

	Traverse(bt)
}
