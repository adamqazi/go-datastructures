package main

import (
	"fmt"
)

type stack struct {
	array []interface{}
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) isEmpty() bool {
	return len(s.array) == 0
}

func (s *stack) peek() interface{} {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return nil
	}

	return s.array[len(s.array)-1]
}

func (s *stack) push(value interface{}) {
	if s.array == nil {
		s.array = make([]interface{}, 0)
	}
	s.array = append(s.array, value)
}

func (s *stack) pop() interface{} {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return nil
	}

	if len(s.array) == 1 {
		value := s.array[0]
		s.array = nil
		return value
	}

	value := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return value
}

func main() {
	fmt.Println("array implementation of stack")

	stk := newStack()
	stk.peek()
	stk.push("google")
	stk.push("udemy")
	stk.push("discord")

	topPop := stk.pop()
	fmt.Println("popped value", topPop)

	topPop1 := stk.pop()
	fmt.Println("popped value", topPop1)

	topPop2 := stk.pop()
	fmt.Println("popped value", topPop2)

	fmt.Println(stk.pop())

	fmt.Println(stk)
}
