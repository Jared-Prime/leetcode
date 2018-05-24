package tree

// MaximumDepth finds the length of the longest path in the tree.
func (root *Node) MaximumDepth() int {
	var left, right, max int

	if root.IsEmpty() {
		return max
	}

	left = root.Left.MaximumDepth()
	right = root.Right.MaximumDepth()

	if left > right {
		max = left
	} else {
		max = right
	}

	return max + 1
}

func (root *Node) IsSymmetric() bool {
	if root.IsEmpty() {
		return true
	}

	return mirror(root.Left, root.Right)
}

func mirror(left, right *Node) bool {
	if left.IsEmpty() && right.IsEmpty() {
		return true
	}

	if !left.IsEmpty() && right.IsEmpty() {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return mirror(left.Right, right.Left) && mirror(left.Left, right.Right)
}
