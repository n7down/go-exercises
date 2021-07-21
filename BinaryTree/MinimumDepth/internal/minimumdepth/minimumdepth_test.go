package minimumdepth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDepthWhereMinimumDepthIsTwo(t *testing.T) {
	expectedDepth := 2

	n4 := &Node{
		Data:  4,
		Left:  nil,
		Right: nil,
	}

	n2 := &Node{
		Data:  2,
		Left:  n4,
		Right: nil,
	}

	n3 := &Node{
		Data:  3,
		Left:  nil,
		Right: nil,
	}

	n1 := &Node{
		Data:  1,
		Left:  n2,
		Right: n3,
	}

	actualDepth := n1.FindMinimumDepth()

	assert.Equal(t, expectedDepth, actualDepth)
}
