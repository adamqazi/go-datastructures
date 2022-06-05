package main

import (
	"fmt"
)

// linkedList represents a doubly linked list.
type linkedList struct {
	head   *node
	tail   *node
	length int
}

// node respresents a node of a doubly linked list.
type node struct {
	value interface{}
	next  *node
	prev  *node
}

// newLinkedList creates a new empty linked list.
func newLinkedList() *linkedList {
	return &linkedList{}
}

// newNode creates a new node with the value provided.
func newNode(value interface{}) *node {
	return &node{value: value}
}

// isEmpty checks whether the list is empty.
func (l *linkedList) isEmpty() bool {
	return l.head == nil && l.tail == nil && l.length == 0
}

// printList prints the entire list along with its length.
func (l *linkedList) printList() {
	if l.isEmpty() {
		fmt.Println("linked list is empty")

		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v <=> ", currentNode.value)
		// uncomment to see if it is working as expected.
		// fmt.Printf("value: %+v, next: %+v, prev: %+v\n", currentNode.value, currentNode.next, currentNode.prev)
		currentNode = currentNode.next
	}

	fmt.Println("\nlength of linked list is", l.length, "\n")
}

// makeAndAssignNode creates a new node with the value provided and assigns it if the list is empty.
func (l *linkedList) makeAndAssignNode(value interface{}) (*node, bool) {
	var assigned bool
	newNode := newNode(value)
	// if linked list is empty, point both head and tail to new node and mark as assigned.
	if l.isEmpty() {
		l.head = newNode
		l.tail = newNode
		l.length += 1
		assigned = true
	}

	return newNode, assigned
}

// append creates a new node with the value provided and inserts it at the end of the list.
func (l *linkedList) append(value interface{}) {
	fmt.Println("=> appending", value, "\n")

	newNode, assigned := l.makeAndAssignNode(value)
	// if new node is not already assigned, append the new node at the end.
	if !assigned {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
		l.length += 1
	}
	// print linked list
	l.printList()
}

// prepend creates a new node with the value provided and inserts it at the beginning of the list.
func (l *linkedList) prepend(value interface{}) {
	fmt.Println("=> prepending", value, "\n")

	newNode, assigned := l.makeAndAssignNode(value)
	// if new node is not already assigned, prepend the new node at the beginning.
	if !assigned {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
		l.length += 1
	}
	// print linked list
	l.printList()
}

// insert creates a new node with the value provided and inserts it at the specified index.
func (l *linkedList) insert(index int, value interface{}) {
	validateIndex(index, l.length)

	fmt.Println("=> inserting", value, "at index", index, "\n")

	newNode, assigned := l.makeAndAssignNode(value)
	// if new node is not already assigned, insert the new node at the specified index.
	if !assigned {
		// if index is head, prepend the new node at the beginning.
		if index == 0 {
			l.prepend(value)

			return
		}
		// if index is after tail, append the new node at the end.
		if index == l.length {
			l.append(value)

			return
		}

		currentNode := l.getNodeAtIndex(index)
		newNode.next = currentNode
		newNode.prev = currentNode.prev
		currentNode.prev.next = newNode
		currentNode.prev = newNode
		l.length += 1
		l.printList()
	}
}

// remove removes the node at the specified index.
func (l *linkedList) remove(index int) {
	validateIndex(index, l.length-1)

	fmt.Println("=> removing node at index", index, "\n")

	if index == 0 { // if index is head, update head to the next node.
		l.head = l.head.next
		l.head.prev = nil
	} else if index == l.length-1 { // if index is tail, update second last node as tail
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		currentNode := l.getNodeAtIndex(index)
		prevNode := currentNode.prev
		prevNode.next = currentNode.next
		currentNode.next.prev = prevNode
	}

	l.length -= 1
	l.printList()
}

// reverse reverses the order of the nodes in the linked list.
func (l *linkedList) reverse() {
	fmt.Println("=> reversing linked list\n")
	// if the linked list is empty or contains only one node, no need to reverse it.
	if l.length <= 1 {
		return
	}

	headNode := l.head
	l.tail = l.head
	currentNode := headNode.next
	for currentNode != nil {
		temp := currentNode.next
		currentNode.next = headNode
		currentNode.prev = temp
		headNode = currentNode
		currentNode = temp
	}
	l.tail.prev = l.head.next
	l.head.next = nil
	l.head = headNode

	l.printList()
}

// getNodeAtIndex returns the node at the specified index.
func (l *linkedList) getNodeAtIndex(index int) *node {
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode
}

// validateIndex validates the provided index against the length of the linked list.
func validateIndex(index, length int) {
	// check if index provided is out of range.
	if index < 0 || index > length {
		panic(fmt.Sprintf("index, %v, out of range; expected index between 0 and %v", index, length))
	}
}

func main() {
	ll := newLinkedList()
	ll.append(10)
	ll.append(5)
	ll.append(16)
	ll.prepend(1)
	ll.insert(2, 99)
	ll.insert(0, 100)
	ll.insert(5, 101)
	ll.insert(7, 200)
	ll.remove(5)
	ll.remove(0)
	ll.remove(5)
	ll.reverse()
	fmt.Printf("head: %+v, tail: %+v, length: %+v", ll.head, ll.tail, ll.length)
	// ll.remove(8)
	// ll.insert(100, 4000)
}
