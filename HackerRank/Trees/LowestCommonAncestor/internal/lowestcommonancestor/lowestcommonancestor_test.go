package lowestcommonancestor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given:
//          1
//        /   \
//       2     3
//      / \   / \
//     4   5 6   7

// LCA(4, 5) = 2
// LCA(4, 6) = 1
// LCA(3, 4) = 1
// LCA(2, 4) = 2

func TestLowestCommonAncestor(t *testing.T) {
	n7 := &Node{
		Data: 7,
	}

	n6 := &Node{
		Data: 6,
	}

	n5 := &Node{
		Data: 5,
	}

	n4 := &Node{
		Data: 4,
	}

	n3 := &Node{
		Data:  3,
		Left:  n6,
		Right: n7,
	}

	n2 := &Node{
		Data:  2,
		Left:  n4,
		Right: n5,
	}

	n1 := &Node{
		Data:  1,
		Left:  n2,
		Right: n3,
	}

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{
			left:     4,
			right:    5,
			expected: 2,
		},
		{
			left:     4,
			right:    6,
			expected: 1,
		},
		{
			left:     3,
			right:    4,
			expected: 1,
		},
		{
			left:     2,
			right:    3,
			expected: 1,
		},
		{
			left:     2,
			right:    6,
			expected: 1,
		},
		{
			left:     2,
			right:    4,
			expected: 2,
		},
	}

	for _, test := range tests {
		actual := n1.FindLowestCommonAncestor(test.left, test.right)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("left: %d, right: %d should return %d", test.left, test.right, test.expected))
	}
}
