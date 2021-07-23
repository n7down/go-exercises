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
	leftFound, _ := left.getLowestCommonAncestor(v1, v2)

	// if leftData != 0 {
	// 	return true, leftData
	// }

	right := n.Right
	rightFound, _ := right.getLowestCommonAncestor(v1, v2)

	// if rightData != 0 {
	// 	return true, rightData
	// }

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
