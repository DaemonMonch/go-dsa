package godsa

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(e interface{}) {
	s.l.PushFront(e)
}

func (s *Stack) Pop() (bool, interface{}) {
	if s == nil || s.l.Len() == 0 {
		return false, nil
	}

	e := s.l.Front()
	s.l.Remove(e)
	return true, e.Value
}

func (s *Stack) Len() int {
	return s.l.Len()
}
