package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	var res *Node
	var err error

	check := assert.New(t)

	for i := 1; i < 5; i++ {
		res, err = newTraversalSubject().Search(i, Preordering)
		res, err = newTraversalSubject().Search(i, Inordering)
		res, err = newTraversalSubject().Search(i, Postordering)
		if i < 4 {
			check.NoError(err)
			check.Equal(i, res.Val)
		} else {
			check.Error(err)
			check.Nil(res)
		}
	}

	res, err = newTraversalSubject().Search(1, 3)
	check.Error(err)
	check.Nil(res)

}
