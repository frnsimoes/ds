package btree

type Node struct {
	left  *Node
	right *Node
	value int
}

func Insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}

	if value < node.value {
		node.left = Insert(node.left, value)
	} else {
		node.right = Insert(node.right, value)
	}

	return node
}

func Search(node *Node, value int) *Node {
	if node == nil || node.value == value {
		return node
	}
	if value < node.value {
		return Search(node.left, value)
	}
	return Search(node.right, value)
}

func Delete(node *Node, value int) *Node {
	return nil
}
