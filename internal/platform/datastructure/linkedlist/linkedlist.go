package linkedlist

import (
	"fmt"
	it "jobs/internal/platform/datastructure/iterator"
	"jobs/internal/platform/datastructure/node"
)

type iterator struct {
	current *node.Node
}

func (it *iterator) Next() *node.Node {
	current := it.current
	it.current = it.current.Next()
	return current
}

func (it *iterator) HasNext() bool {
	return it.current != nil
}

type LinkedList struct {
	head    *node.Node
	current *node.Node
	length  int
}

func (li *LinkedList) Iterator() it.IIterator {
	head := li.head
	return &iterator{head}
}

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (li *LinkedList) Head() *node.Node {
	return li.head
}

// Prints elements from the list
func (li *LinkedList) Print() {
	head := li.head
	for head != nil {
		fmt.Printf("%v\n", head.Item())
		head = head.Next()
	}
}

// Inserts at the end of list
func (li *LinkedList) Append(item int, status, jobType string) *node.Node {
	newNode := node.New(item, status, jobType) // new Node
	if li.head == nil {
		li.head = newNode
		li.current = newNode
		li.length = 1
		return newNode
	}

	li.current.SetNext(newNode)
	li.current = newNode
	li.length++
	return newNode
}

// Removes an element from front of the list
func (li *LinkedList) RemoveFront() *node.Node {
	head := li.head
	if head != nil {
		nextNode := head.Next()
		li.head = nextNode
		li.length--
		head.SetNext(nil)
		return head
	}

	return nil
}
