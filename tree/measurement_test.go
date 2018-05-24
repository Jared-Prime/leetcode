package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newUnbalancedTree() *Node {
	return &Node{
		Left: &Node{
			Right: &Node{},
		},
		Right: &Node{},
	}
}

func newBalancedTree() *Node {
	return &Node{
		Left:  &Node{},
		Right: &Node{},
	}
}

func newSymmetricTree() *Node {
	return &Node{
		Val: 0,
		Left: &Node{
			Val:   1,
			Left:  &Node{Val: 2},
			Right: &Node{Val: 3},
		},
		Right: &Node{
			Val:   1,
			Left:  &Node{Val: 3},
			Right: &Node{Val: 2},
		},
	}
}

func newAssymetricTree() *Node {
	return &Node{
		Val: 0,
		Left: &Node{
			Val:   1,
			Left:  &Node{Val: 2},
			Right: &Node{Val: 2},
		},
		Right: &Node{
			Val:   1,
			Right: &Node{Val: 2},
		},
	}
}

func anotherAssymmetricTree() *Node {
	return &Node{
		Val:  0,
		Left: &Node{},
		Right: &Node{
			Val:   1,
			Left:  &Node{Val: 2},
			Right: &Node{Val: 3},
		},
	}
}

func yetAnotherAssymetricTree() *Node {
	return &Node{
		Val: 0,
		Left: &Node{
			Val:   1,
			Left:  &Node{Val: 2},
			Right: &Node{Val: 2},
		},
		Right: &Node{
			Val:  1,
			Left: &Node{Val: 2},
		},
	}
}

func newEmptyTree() *Node {
	return &Node{}
}

func TestMaximumDepth(t *testing.T) {
	assert.Equal(t, 3, newUnbalancedTree().MaximumDepth())
	assert.Equal(t, 2, newBalancedTree().MaximumDepth())
}

func TestIsSymmetric(t *testing.T) {
	var missing *Node
	check := assert.New(t)

	check.True(missing.IsSymmetric())
	check.True(newEmptyTree().IsSymmetric())
	check.True(newBalancedTree().IsSymmetric())
	check.True(newSymmetricTree().IsSymmetric())
	check.False(newUnbalancedTree().IsSymmetric())
	check.False(newAssymetricTree().IsSymmetric())
	check.False(anotherAssymmetricTree().IsSymmetric())
	check.False(yetAnotherAssymetricTree().IsSymmetric())
}
