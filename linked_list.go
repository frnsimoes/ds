package main

import "errors"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data int) {
	node := &Node{data: data}
	if l.IsEmpty() {
		l.head = node
	} else {
		var lastNode *Node
		// this is just bad, gotta implement tail
		for node := l.head; node != nil; node = node.next {
			if node.next == nil {
				lastNode = node
			}
		}
		lastNode.next = node
	}
}

func (l *LinkedList) Add(data int) {
	node := &Node{data: data}

	node.next = l.head
	l.head = node
}

func (l *LinkedList) Insert(index int, data int) error {
	node := &Node{data: data}
	foundNode := l.Search(index)

	if index > l.Size() {
		return errors.New("Index out of range")
	}
	if index == 0 {
		l.head = node
	} else {
		previousNode := l.Search(index - 1)
		previousNode.next = node
		node.next = foundNode
	}

	return nil
}

func (l *LinkedList) Remove(index int) error {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		l.head = l.head.next
		return nil
	}

	previousNode := l.Search(index - 1)
	nodeToRemove := previousNode.next

	// PreviousNode doesn't point to nodeToRemove.
	previousNode.next = nodeToRemove.next

	// nodeToRemove doesn't point to anything. Garbage collector will take care of it.
	nodeToRemove.next = nil

	return nil
}

func (l *LinkedList) Search(index int) *Node {
	var count int = 0
	for node := l.head; node != nil; node = node.next {
		if count == index {
			return node
		}
		count++
	}
	return nil
}

func (l *LinkedList) Size() int {
	var count int = 0
	for node := l.head; node != nil; node = node.next {
		count++
	}
	return count
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}
