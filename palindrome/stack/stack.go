package stack

type Stack struct {
	items []rune
}

func New() *Stack {
	return &Stack{
		items: make([]rune, 0),
	}
}

func (s *Stack) Values() []rune {
	return s.items
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		panic("stack underflow")
	}
	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popped
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
