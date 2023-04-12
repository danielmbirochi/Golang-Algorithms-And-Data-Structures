package main

import (
	"fmt"

	"github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/palindrome/stack"
)

type Node struct {
	Val  rune
	Next *Node
}

func IsPalindrome(n *Node) bool {
	curr := n
	runner := n
	s := stack.New()

	for runner != nil && runner.Next != nil { // fill the stack w/ the first half of the list.
		s.Push(curr.Val)
		curr = curr.Next
		runner = runner.Next.Next
	}

	if runner != nil { // Odd list.
		curr = curr.Next // Set curr in the second half of the list.
	}

	for curr != nil {
		if s.Pop() != curr.Val {
			return false
		}
		curr = curr.Next
	}

	return true
}

func ToLinkedList(s string) *Node {
	var head *Node
	for _, char := range s {
		curr := &Node{
			Val:  char,
			Next: head,
		}
		head = curr
	}
	return head
}

func main() {
	words := []string{"anna", "civic", "madam", "level", "mom", "racecar", "tenet", "mygym", "topspot", "nolemon nomelon", "casa", "lama", "tenis"}

	for _, word := range words {
		fmt.Printf("%s is palindrome? %t\n", word, IsPalindrome(ToLinkedList(word)))
	}
}
