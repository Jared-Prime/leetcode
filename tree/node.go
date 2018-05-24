package tree

import "fmt"

// Node defines the nodes ( both root and children ) for a binary tree
type Node struct {
	Val         int
	Left, Right *Node
}

func (node *Node) String() string {
	if node.IsLeaf() {
		return fmt.Sprintf("%d", node.Val)
	}

	return fmt.Sprintf("%d->[ %v %v ]", node.Val, node.Left, node.Right)
}

// IsLeaf returns whether or not a node has children. A node without children is called a leaf
func (node Node) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

// IsEmpty checks if the pointer is nil
func (node *Node) IsEmpty() bool {
	return node == nil
}
