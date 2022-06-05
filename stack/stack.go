package main

import (
	"fmt"
)

// stack represents the stack data structure.
type stack struct {
	top    *node
	bottom *node
	length int
}

// node represents a node of a stack.
type node struct {
	value interface{}
	next  *node
}

// newStack creates a new empty stack.
func newStack() *stack {
	return &stack{}
}

// newNode creates a new node with the value provided.
func newNode(value interface{}) *node {
	return &node{value: value}
}

// isEmpty checks whether the stack is empty.
func (s *stack) isEmpty() bool {
	return s.top == nil && s.bottom == nil && s.length == 0
}

// printStack prints the entire stack along with its length.
func (s *stack) printStack() {
	if s.isEmpty() {
		fmt.Println("stack is empty")
	}

	currentNode := s.top
	for currentNode != nil {
		fmt.Printf("%v\n↓\n", currentNode.value)
		// uncomment to see if it is working as expected.
		// fmt.Printf("value: %+v, next: %+v\n", currentNode.value, currentNode.next)
		currentNode = currentNode.next
	}

	fmt.Println("\nlength of stack is", s.length, "\n")
}

// makeAndAssignNode creates a new node with the value provided and assigns it if the stack is empty.
func (s *stack) makeAndAssignNode(value interface{}) (*node, bool) {
	var assigned bool
	newNode := newNode(value)
	// if stack is empty, point both top and bottom to new node and mark as assigned.
	if s.isEmpty() {
		s.top = newNode
		s.bottom = newNode
		s.length += 1
		assigned = true
	}

	return newNode, assigned
}

// peek returns the node at the top of the stack.
func (s *stack) peek() *node {
	return s.top
}

// push creates a new node with the value provided and pushes it at the top of the stack.
func (s *stack) push(value interface{}) {
	fmt.Println("=> pushing", value, "\n")

	newNode, assigned := s.makeAndAssignNode(value)
	if !assigned {
		newNode.next = s.top
		s.top = newNode
		s.length += 1
	}
	// print stack
	s.printStack()
}

// pop removes the top node from the stack and returns it.
func (s *stack) pop() *node {
	if s.isEmpty() {
		fmt.Println("stack is empty, returning nil")

		return nil
	}

	if s.top == s.bottom {
		s.bottom = nil
	}

	topNode := s.top
	s.top = topNode.next
	topNode.next = nil
	s.length -= 1

	return topNode
}

func main() {
	fmt.Println("linked list implementation of stack")

	stk := newStack()
	stk.push("google")
	stk.push("udemy")
	stk.push("discord")

	topPeek := stk.peek()
	fmt.Printf("=> peeking top node, value: %+v, next: %+v\n", topPeek.value, topPeek.next)

	topPop := stk.pop()
	fmt.Printf("=> popping top node, value: %+v, next: %+v\n", topPop.value, topPop.next)

	stk.printStack()

	topPop2 := stk.pop()
	fmt.Printf("=> popping top node, value: %+v, next: %+v\n", topPop2.value, topPop2.next)

	stk.printStack()

	topPop3 := stk.pop()
	fmt.Printf("=> popping top node, value: %+v, next: %+v\n", topPop3.value, topPop3.next)

	fmt.Printf("current state of the stack, %+v\n", stk)
	stk.printStack()
}
