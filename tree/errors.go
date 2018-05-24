package tree

import (
	"fmt"
)

// UnknownOrdering may be used to signal to users that the given integer does not match a traversal order
func UnknownOrdering(given int) error {
	return fmt.Errorf("unkown ordering: %d", given)
}

// NotFound may be used to signal to users that the target value cannot be found in the tree
func NotFound(target int) error {
	return fmt.Errorf("could not find node matching: %d", target)
}
