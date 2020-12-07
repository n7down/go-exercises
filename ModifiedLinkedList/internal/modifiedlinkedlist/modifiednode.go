package modifiednode

import "fmt"

type Node struct {
	nextNode *Node

	// max length 5
	data []int
}

// {3, 9, 7, 5, 6} -> {1, 4, 8}
func (n *Node) Get(d int) int {
	if d < 5 {
		return n.data[d]
	}

	// d - 5 = 2
	i := d - 5
	currentNode := n.nextNode
	for currentNode != nil {
		if i > 5 {
			currentNode = currentNode.nextNode
			i = i - 5
		}
	}
	return currentNode.data[i]
}

func (n *Node) Insert(d int) {
	if len(n.data) < 5 {
		n.data = append(n.data, d)
		return
	}

	currentNode := n.nextNode
	prevNode := n
	for currentNode != nil {
		currentNode = currentNode.nextNode
		prevNode = currentNode
	}

	if len(prevNode.data) < 5 {
		prevNode.data = append(prevNode.data, d)
	} else {
		n := &Node{}
		n.data = append(n.data, d)
		prevNode.nextNode = n
	}
}

func (n *Node) Print() {
	fmt.Println(fmt.Sprintf("%v", n.data))
	currentNode := n.nextNode
	for currentNode != nil {
		fmt.Println(fmt.Sprintf("%v", n.data))
	}
}
