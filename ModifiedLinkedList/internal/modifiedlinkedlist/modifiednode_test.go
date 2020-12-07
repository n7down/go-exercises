package modifiednode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeWithFiveElements(t *testing.T) {
	expected := 6
	n := &Node{}
	n.Insert(3)
	n.Insert(9)
	n.Insert(7)
	n.Insert(5)
	n.Insert(6)
	actual := n.Get(4)
	assert.Equal(t, expected, actual)
}

func TestNodeWithSixElements(t *testing.T) {
	expected := 1
	n := &Node{}
	n.Insert(3)
	n.Insert(9)
	n.Insert(7)
	n.Insert(5)
	n.Insert(6)
	n.Insert(1)
	actual := n.Get(5)
	assert.Equal(t, expected, actual)
}
