package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	items []int
	size  int
	sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		items: nil,
		size:  0,
	}
}

func (s *Stack) Push(t int) {
	s.Lock()
	s.items = append(s.items, t)
	s.size++
	s.Unlock()
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("栈为空")
	}
	s.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.Unlock()
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
