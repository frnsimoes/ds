package btree

import (
	"testing"
)

func TestSearch(t *testing.T) {
	root := &Node{value: 10}

	root.left = &Node{value: 5}
	root.right = &Node{value: 15}

	// Test search for existing value
	result := Search(root, 15)
	if result == nil || result.value != 15 {
		t.Errorf("Expected node with value 15, but got %v", result)
	}

	// Test search for non-existing value
	result = Search(root, 20)
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}

func TestInsertion(t *testing.T) {
	root := &Node{value: 10}

	Insert(root, 5)
	Insert(root, 15)

	if root.left.value != 5 {
		t.Errorf("Expected left node to have value 5, but got %v", root.left.value)
	}

	if root.right.value != 15 {
		t.Errorf("Expected right node to have value 15, but got %v", root.right.value)
	}

}
