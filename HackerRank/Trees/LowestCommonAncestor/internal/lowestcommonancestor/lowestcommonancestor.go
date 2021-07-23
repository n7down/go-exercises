package lowestcommonancestor

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) getLowestCommonAncestor(v1, v2 int) (bool, int) {
	if n == nil {
		return false, 0
	}

	left := n.Left
	leftFound, leftData := left.getLowestCommonAncestor(v1, v2)

	if leftData != 0 {
		return true, leftData
	}

	right := n.Right
	rightFound, rightData := right.getLowestCommonAncestor(v1, v2)

	if rightData != 0 {
		return true, rightData
	}

	if leftFound && n.Data == v1 {
		return true, n.Data
	}

	if leftFound && n.Data == v2 {
		return true, n.Data
	}

	if rightFound && n.Data == v1 {
		return true, n.Data
	}

	if rightFound && n.Data == v2 {
		return true, n.Data
	}

	if n.Data == v1 {
		return true, 0
	}

	if n.Data == v2 {
		return true, 0
	}

	if leftFound && rightFound {
		return true, n.Data
	}

	if leftFound {
		return true, 0
	}

	if rightFound {
		return true, 0
	}

	return false, 0
}

func (n *Node) FindLowestCommonAncestor(v1, v2 int) int {
	found, commonAncestor := n.getLowestCommonAncestor(v1, v2)
	if found {
		return commonAncestor
	}

	return 0
}
