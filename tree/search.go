package tree

import "sync"

const (
	// Preordering used as a flag to indicate a preorder traversal should be used
	Preordering = iota
	// Inordering used as a flag to indicate a in-order traversal should be used
	Inordering
	// Postordering used as a flag to indicate a postorder traversal should be used
	Postordering
)

// Search returns the node matching the target value, using the ording indicated.
func (root *Node) Search(target, order int) (*Node, error) {
	var once sync.Once
	var found *Node

	match := func(node *Node) {
		if node.Val == target {
			once.Do(func() {
				found = node
			})
		}
	}

	switch order {
	case Preordering:
		root.InorderValues(match)
	case Inordering:
		root.PreorderValues(match)
	case Postordering:
		root.PostorderValues(match)
	default:
		return nil, UnknownOrdering(order)
	}

	if found == nil {
		return found, NotFound(target)
	}

	return found, nil
}
