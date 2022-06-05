package main

import (
	"fmt"
)

type queue struct {
	first  *node
	last   *node
	length int
}

type node struct {
	value interface{}
	next  *node
}

func newQueue() *queue {
	return &queue{}
}

func newNode(value interface{}) *node {
	return &node{value: value}
}

func (q *queue) isEmpty() bool {
	return q.first == nil && q.last == nil && q.length == 0
}

func (q *queue) print() {
	if q.isEmpty() {
		fmt.Println("queue is empty")
	}

	currentNode := q.first
	for currentNode != nil {
		fmt.Printf("%v -> ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println("\nlenght of the queue is", q.length, "\n")
}

func (q *queue) makeAndAssignNode(value interface{}) (*node, bool) {
	newNode := newNode(value)

	var assigned bool
	if q.isEmpty() {
		q.first = newNode
		q.last = newNode
		q.length += 1
		assigned = true
	}

	return newNode, assigned
}

func (q *queue) peek() *node {
	fmt.Println("=> peeking queue")

	if q.isEmpty() {
		fmt.Println("queue is empty")
	}

	return q.first
}

func (q *queue) enqueue(value interface{}) {
	fmt.Println("=> enqueueing", value)

	newNode, assigned := q.makeAndAssignNode(value)
	if !assigned {
		q.last.next = newNode
		q.last = newNode
		q.length += 1
	}

	q.print()
}

func (q *queue) dequeue() *node {
	fmt.Println("=> dequeueing", q.first.value)

	if q.isEmpty() {
		fmt.Println("queue is empty")

		return nil
	}

	currentNode := q.first

	if q.first == q.last && q.length == 1 {
		q.first = nil
		q.last = nil
		q.length -= 1
	} else {
		q.first = currentNode.next
		currentNode.next = nil
		q.length -= 1
	}

	q.print()

	return currentNode
}

func main() {
	fmt.Println("linked list implementation of queues\n")

	q := newQueue()

	q.enqueue("Joy")

	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	q.enqueue("Matt")

	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	q.enqueue("Pavel")

	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	q.enqueue("Sameer")

	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	fmt.Println(q.peek().value, "\n")

	dq := q.dequeue()
	fmt.Println(dq.value, dq.next)
	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	dq1 := q.dequeue()
	fmt.Println(dq1.value, dq1.next)
	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	dq2 := q.dequeue()
	fmt.Println(dq2.value, dq2.next)
	fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")

	dq3 := q.dequeue()
	fmt.Println(dq3.value, dq3.next)
	// fmt.Println(q.first.value, q.first.next, q.last.value, q.last.next, q.length, "\n")
}
