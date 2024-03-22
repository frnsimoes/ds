package btree

// Binary Search Tree, actually
type BinaryTree struct {
	root *Node
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func insert(btree *BinaryTree, node *Node) *Node {
	// insert a node into a binary tree
	_ = btree
	_ = node
	return nil
}

type traverseFunc func(node *Node, value int) *Node

func search(node *Node, value int) *Node {
	if value == node.value {
		return node
	}

	if node == nil {
		return nil
	}

	if value < node.value {
		return search(node.left, value)
	} else {
		return search(node.right, value)
	}
}
