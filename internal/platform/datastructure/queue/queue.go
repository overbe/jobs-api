package queue

import (
	it "jobs/internal/platform/datastructure/iterator"
	"jobs/internal/platform/datastructure/linkedlist"
	"jobs/internal/platform/datastructure/node"
)

type IQueue interface {
	Enqueue(item int) *node.Node
	Dequeue() *node.Node
	Peek() *node.Node
	Iterator() it.IIterator
}

type Queue struct {
	li *linkedlist.LinkedList
}

func New() *Queue {
	list := linkedlist.New()
	return &Queue{list}
}

func (q *Queue) Iterator() it.IIterator {
	return q.li.Iterator()
}

// Top of the list
func (q *Queue) Peek() *node.Node {
	return q.li.Head()
}

// Enqueues element
func (q *Queue) Enqueue(item int, status, jobType string) *node.Node {
	newNode := q.li.Append(item, status, jobType)
	return newNode
}

// Dequeue element
func (q *Queue) Dequeue() *node.Node {
	removedNode := q.li.RemoveFront()
	return removedNode
}
