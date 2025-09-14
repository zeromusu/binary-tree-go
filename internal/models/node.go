package models

type Node struct {
	key        int
	leftChild  *Node
	rightChild *Node
}

func createNode(key int) *Node {
	return &Node{key: key}
}

func (n *Node) removeNode() bool {
	if n == nil {
		return false
	}

	n.setLeftChild(nil)
	n.setRightChild(nil)

	return true
}

func (n *Node) searchNode(key int) *Node {
	if n == nil {
		return nil
	}
	if n.getKey() == key {
		return n
	}

	if n.getKey() > key {
		return n.getLeftChild().searchNode(key)
	}

	return n.getRightChild().searchNode(key)
}

func (n *Node) getKey() int {
	return n.key
}

func (n *Node) setKey(key int) {
	n.key = key
}

func (n *Node) getLeftChild() *Node {
	return n.leftChild
}

func (n *Node) setLeftChild(node *Node) {
	n.leftChild = node
}

func (n *Node) getRightChild() *Node {
	return n.rightChild
}

func (n *Node) setRightChild(node *Node) {
	n.rightChild = node
}
