package main

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

func (l *LinkedList) Insert(index int, data int) {
	node := &Node{data: data}
	rNode := l.Search(index)

}

func (l *LinkedList) Remove(index int) {}

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
