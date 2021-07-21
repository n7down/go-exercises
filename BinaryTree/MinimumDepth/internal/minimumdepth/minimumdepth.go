package minimumdepth

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) FindMinimumDepth() int {
	if n == nil {
		return 0
	}

	leftDepth := n.Left.FindMinimumDepth()
	if n.Left == nil {
		return 1 + leftDepth
	}

	rightDepth := n.Right.FindMinimumDepth()
	if n.Right == nil {
		return 1 + rightDepth
	}

	return min(leftDepth, rightDepth) + 1
}
