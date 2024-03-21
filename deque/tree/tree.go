package tree

type Node struct {
	data     int
	children []*Node
}

type Tree struct {
	root *Node
}
