package btree

import (
	"testing"
)

func TestSearch(t *testing.T) {
	root := &Node{value: 10}
	_ = BinaryTree{root: root}

	root.left = &Node{value: 5}
	root.right = &Node{value: 15}

	// Test search for existing value
	result := search(root, 15)
	if result == nil || result.value != 15 {
		t.Errorf("Expected node with value 15, but got %v", result)
	}

	// Test search for non-existing value
	result = search(root, 20)
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
