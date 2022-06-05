package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type node struct {
	value int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func newNode(value int) *node {
	return &node{value: value}
}

func newBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

func (b *binarySearchTree) makeAndAssignNode(value int) (*node, bool) {
	newNode := newNode(value)

	var assigned bool
	if b.root == nil {
		b.root = newNode
		assigned = true
	}

	return newNode, assigned
}

func (b *binarySearchTree) insert(value int) {
	newNode, assigned := b.makeAndAssignNode(value)
	if !assigned {
		currentNode := b.root
		for {
			if value < currentNode.value { // go left
				if currentNode.left != nil {
					currentNode = currentNode.left
				} else {
					currentNode.left = newNode
					break
				}
			} else { // go right
				if currentNode.right != nil {
					currentNode = currentNode.right
				} else {
					currentNode.right = newNode
					break
				}
			}
		}
	}
}

func (b *binarySearchTree) lookup(value int) *node {
	if b.root == nil {
		fmt.Println("binary search tree is empty")

		return nil
	}

	currentNode := b.root
	for currentNode != nil {
		if value == currentNode.value {
			return currentNode
		} else if value < currentNode.value { // go left
			currentNode = currentNode.left
		} else { // go right
			currentNode = currentNode.right
		}
	}

	return nil
}

// func (b *binarySearchTree) remove(value int) {
// 	if b.root == nil {
// 		fmt.Println("binary search tree is empty")
// 		return
// 	}

// 	var parentNode *node
// 	currentNode := b.root
// 	for {
// 		if value < currentNode.value { // go left
// 			parentNode = currentNode
// 			currentNode = currentNode.left
// 		} else if value > currentNode.value { // go right
// 			parentNode = currentNode
// 			currentNode = currentNode.right
// 		} else if value == currentNode.value {
// 			if currentNode.right == nil {
// 				if parentNode == nil {
// 					b.root = currentNode.left
// 				} else {

// 				}
// 			}
// 		}
// 	}
// }

func main() {
	bst := newBinarySearchTree()
	bst.insert(9)
	bst.insert(4)
	bst.insert(6)
	bst.insert(20)
	bst.insert(170)
	bst.insert(15)
	bst.insert(1)
	spew.Dump(bst)
	spew.Dump(bst.lookup(170))
	spew.Dump(bst.lookup(6))
	spew.Dump(bst.lookup(20))
	spew.Dump(bst.lookup(90))
}

//      9
//   4     20
// 1  6  15  170
