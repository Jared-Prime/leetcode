package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeStringer(t *testing.T) {
	assert.Equal(t, "2->[ 1 3 ]", fmt.Sprintf("%v", newTraversalSubject()))
}
