package utils

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	data []T
	head int
	lock sync.Mutex
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, capacity),
	}
}

func (s *Stack[T]) Count() int {
	s.lock.Lock()
	count := s.head
	s.lock.Unlock()
	return count
}

func (s *Stack[T]) Capacity() int {
	return len(s.data)
}

func (s *Stack[T]) Push(val T) error {

	s.lock.Lock()
	if len(s.data) == s.head {
		s.lock.Unlock()
		return errors.New("the buffer is full and cannot take more data")
	}

	s.data[s.head] = val
	s.head += 1
	s.lock.Unlock()

	return nil
}

func (s *Stack[T]) Peek() (T, bool) {
	var found bool
	var v T
	s.lock.Lock()
	if s.head > 0 {
		idx := s.head - 1
		v = s.data[idx]
		found = true
	}

	s.lock.Unlock()
	return v, found
}

func (s *Stack[T]) Pop() (T, bool) {

	var found bool
	var v T
	var blank T
	s.lock.Lock()
	if s.head > 0 {
		idx := s.head - 1
		v = s.data[idx]
		found = true
		s.data[idx] = blank
		s.head -= 1
	}

	s.lock.Unlock()
	return v, found
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == 0
}
