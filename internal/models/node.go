package models

type Node struct {
	Key        int   `json:"key"`
	Parent     *Node `json:"parent"`
	LeftChild  *Node `json:"left_child"`
	RightChild *Node `json:"right_child"`
}

func CreateNode(key int, parent *Node, left *Node, right *Node) *Node {
	node := &Node{Key: key}
	if parent != nil {
		node.SetParent(parent)
	}
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

	n.SetParent(nil)
	n.SetLeftChild(nil)
	n.SetRightChild(nil)

	return true
}

func (n *Node) GetKey() int {
	return n.Key
}

func (n *Node) SetKey(key int) {
	n.Key = key
}

func (n *Node) GetParent() *Node {
	return n.Parent
}

func (n *Node) SetParent(node *Node) {
	n.Parent = node
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
