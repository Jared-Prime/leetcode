package tree

// InorderValues returns the values of each node in the tree in order of left < root < right
func (root *Node) InorderValues(actions ...func(*Node)) []int {
	var values []int

	left, right := root.Left, root.Right

	if left != nil {
		values = append(values, left.InorderValues(actions...)...)
	}

	values = append(values, process(root, actions...))

	if right != nil {
		values = append(values, right.InorderValues(actions...)...)
	}

	return values
}

// PreorderValues returns the values of each node in the tree in order of root < left < right
func (root *Node) PreorderValues(actions ...func(*Node)) []int {
	var values []int

	left, right := root.Left, root.Right

	values = append(values, process(root, actions...))

	if left != nil {
		values = append(values, left.InorderValues(actions...)...)
	}

	if right != nil {
		values = append(values, right.InorderValues(actions...)...)
	}

	return values
}

// PostorderValues returns the values of each node in the tree in order of left < right < root
func (root *Node) PostorderValues(actions ...func(*Node)) []int {
	var values []int

	left, right := root.Left, root.Right

	if left != nil {
		values = append(values, left.InorderValues(actions...)...)
	}

	if right != nil {
		values = append(values, right.InorderValues(actions...)...)
	}

	values = append(values, process(root, actions...))

	return values
}

func process(node *Node, actions ...func(*Node)) int {
	for _, actOn := range actions {
		actOn(node)
	}

	return node.Val
}
