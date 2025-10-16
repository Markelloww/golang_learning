package main

import (
	"errors"
	"fmt"
)

type stack[T any] struct {
	data []T
}

func (s *stack[T]) push(val T) {
	s.data = append(s.data, val)
}

func (s *stack[T]) pop() (T, error) {
	if len(s.data) == 0 {
		var empty T
		return empty, errors.New("stack is already empty")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

func main() {
	var s stack[int]

	s.push(10)
	s.push(20)

	val, _ := s.pop()
	fmt.Println(val)

	val, _ = s.pop()
	fmt.Println(val)

	val, err := s.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}
