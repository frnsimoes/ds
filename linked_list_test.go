package main

import "testing"

func TestSize(t *testing.T) {
	l := LinkedList{}
	l.Append(1)
	l.Append(50)

	got := l.Size()
	expected := 2

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestAppend(t *testing.T) {
	l := LinkedList{}
	l.Append(1)
	l.Append(50)

	expected := 1
	got := l.head.data
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)

	}

	expected = 50
	got = l.head.next.data
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestAdd(t *testing.T) {
	l := LinkedList{}

	// text nodes.next change
	l.Add(1)
	l.Add(50)

	expectedData := 50
	got := l.head.data
	if got != expectedData {
		t.Errorf("Expected %d, got %d", expectedData, got)
	}

	// test next node
	expectedData = 1
	got = l.head.next.data
	if got != expectedData {
		t.Errorf("Expected %d, got %d", expectedData, got)
	}

	var expectedNode *Node = nil
	var gotNode *Node = l.head.next.next
	if got != expectedData {
		t.Errorf("Expected %v, got %v", expectedNode, gotNode)
	}
}

func TestSearch(t *testing.T) {
	l := LinkedList{}
	l.Add(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	found := l.Search(2).data
	expected := 2
	if found != expected {
		t.Errorf("Expected %v, got %v", expected, found)
	}

}

func TestInsert(t *testing.T) {
	l := LinkedList{}
	l.Add(0)
	l.Append(1)
	l.Append(3)
	l.Append(4)

	l.Insert(2, 2)

	expected := 2
	got := l.Search(2).data

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}

}

func TestRemove(t *testing.T) {
	l := LinkedList{}
	l.Add(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Remove(2)

	expected := 3
	got := l.Search(2).data

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
