package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func (tree *BST) FindClosestValue(target int) int {
	currentNode := tree
	closestNode := tree
	if closestNode.Value == target {
		return closestNode.Value
	}

	for currentNode != nil {
		if calculateAbsoluteDiff(closestNode.Value, target) > calculateAbsoluteDiff(currentNode.Value, target) {
			closestNode = currentNode
		}
		if closestNode.Value == target {
			return closestNode.Value
		}
		if currentNode.Value > target {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	return closestNode.Value
}

func calculateAbsoluteDiff(x, y int) int {
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
	target := -7
	tree := &BST{
		Value: 10,
		Left: &BST{
			Value: -5,
			Right: &BST{
				Value: 5,
			},
			Left: &BST{
				Value: -20,
				Left: &BST{
					Value: 1,
				},
			},
		},
		Right: &BST{
			Value: 11,
			Right: &BST{
				Value: 22,
			},
			Left: &BST{
				Value: 13,
				Right: &BST{
					Value: 14,
				},
			},
		},
	}

	fmt.Println(tree.FindClosestValue(target))
}
