package main

import (
	"errors"
	"fmt"
)

type StackNode[T any] struct {
	data T
	next *StackNode[T]
}

type Stack[T any] struct {
	top *StackNode[T]
}

func (s *Stack[T]) pop() (T, error) {
	if s.top == nil {
		var zero T
		return zero, errors.New("stack is empty")
	}
	temData := s.top.data
	s.top = s.top.next
	return temData, nil
}

func (s *Stack[T]) Push(data T) {
	newNode := &StackNode[T]{data: data, next: s.top}
	s.top = newNode
}

func (s *Stack[T]) peek() T {
	if s.top == nil {
		var zero T
		return zero
	}

	return s.top.data
}

func (s *Stack[T]) isEmpty() bool {
	if s.top == nil {
		return true
	}

	return false
}

func main() {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	val, err := stack.pop()
	if err == nil {
		fmt.Println("Popped:", val)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println(stack.top)
}
