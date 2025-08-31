package models

type Node struct {
	Key        int   `json:"key"`
	LeftChild  *Node `json:"left_child"`
	RightChild *Node `json:"right_child"`
}

func CreateNode(key int, left *Node, right *Node) *Node {
	node := &Node{Key: key}
	if left != nil {
		node.SetLeftChild(left)
	}
	if right != nil {
		node.SetRightChild(right)
	}
	return node
}

func (n *Node) DelNode() bool {
	if n == nil {
		return false
	}

	n.SetLeftChild(nil)
	n.SetRightChild(nil)

	return true
}

func (n *Node) GetKey() int {
	return n.Key
}

func (n *Node) GetLeftChild() *Node {
	return n.LeftChild
}

func (n *Node) SetLeftChild(node *Node) {
	n.LeftChild = node
}

func (n *Node) GetRightChild() *Node {
	return n.RightChild
}

func (n *Node) SetRightChild(node *Node) {
	n.RightChild = node
}
