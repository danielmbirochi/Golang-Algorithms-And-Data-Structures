package main

import (
	"fmt"

	"github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/stack-from-queue/queue"
)

type Stack struct {
	queue *queue.Queue
}

func New() *Stack {
	return &Stack{
		queue: queue.New(),
	}
}

func (s *Stack) Push(x int) {
	temp := queue.New()
	temp.Enqueue(x)
	for !s.queue.IsEmpty() {
		temp.Enqueue(s.queue.Dequeue())
	}
	s.queue = temp
}

func (s *Stack) Pop() int {
	if s.queue.IsEmpty() {
		panic("stack underflow")
	}
	p := s.queue.Dequeue()
	fmt.Println("Popped element : ", p)
	return p
}

func main() {
	s := New()
	s.Push(1)
	s.Push(3)
	s.Push(5)
	s.Pop()
	s.Push(6)
	s.Pop()
	s.Pop()
	s.Pop()
}
