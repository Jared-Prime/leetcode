package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTraversalSubject() *Node {
	return &Node{
		Val:   2,
		Left:  &Node{Val: 1},
		Right: &Node{Val: 3},
	}
}

func incVal(n *Node) { n.Val++ }

func TestInorderValues(t *testing.T) {
	check := assert.New(t)

	check.Equal([]int{1, 2, 3}, newTraversalSubject().InorderValues())

	check.Equal([]int{2, 3, 4}, newTraversalSubject().InorderValues(incVal))
}

func TestPostorderValues(t *testing.T) {
	check := assert.New(t)

	check.Equal([]int{1, 3, 2}, newTraversalSubject().PostorderValues())

	check.Equal([]int{2, 4, 3}, newTraversalSubject().PostorderValues(incVal))
}

func TestPreorderValues(t *testing.T) {
	check := assert.New(t)

	check.Equal([]int{2, 1, 3}, newTraversalSubject().PreorderValues())

	check.Equal([]int{3, 2, 4}, newTraversalSubject().PreorderValues(incVal))
}
