package models

import (
	"fmt"
)

var root *Node

func addNode(key int) error {
	if root == nil {
		root = createNode(key)
		return nil
	}
	node := root.searchNode(key)
	if node != nil {
		return fmt.Errorf("%d is already inserted", key)
	}

	node = createNode(key)
	cursor := root
	for {
		if cursor.getKey() > key {
			if cursor.getLeftChild() == nil {
				cursor.setLeftChild(node)
				break
			}
			cursor = cursor.getLeftChild()
		} else {
			if cursor.getRightChild() == nil {
				cursor.setRightChild(node)
				break
			}
			cursor = cursor.getRightChild()
		}
	}
	root = rebalance(root)
	return nil
}

func deleteNode(key int) error {
	if root == nil {
		return fmt.Errorf("tree is empty")
	}

	var deleted bool
	root, deleted = deleteNodeRec(root, key)
	if !deleted {
		return fmt.Errorf("%d not found", key)
	}

	root = rebalance(root)
	return nil
}

func deleteNodeRec(node *Node, key int) (*Node, bool) {
	if node == nil {
		return nil, false
	}

	if key < node.getKey() {
		left, deleted := deleteNodeRec(node.getLeftChild(), key)
		node.setLeftChild(left)
		return node, deleted
	} else if key > node.getKey() {
		right, deleted := deleteNodeRec(node.getRightChild(), key)
		node.setRightChild(right)
		return node, deleted
	} else {
		if node.getLeftChild() == nil {
			return node.getRightChild(), true
		} else if node.getRightChild() == nil {
			return node.getLeftChild(), true
		}

		minRight := node.getRightChild()
		for minRight.getLeftChild() != nil {
			minRight = minRight.getLeftChild()
		}
		node.setKey(minRight.getKey())
		right, _ := deleteNodeRec(node.getRightChild(), minRight.getKey())
		node.setRightChild(right)

		return node, true
	}
}

func rebalance(node *Node) *Node {
	if node == nil {
		return nil
	}

	node.setLeftChild(rebalance(node.getLeftChild()))
	node.setRightChild(rebalance(node.getRightChild()))

	balance := getBalance(node)

	if balance > 1 {
		if getBalance(node.getLeftChild()) < 0 {
			node.setLeftChild(rotateLeft(node.getLeftChild()))
		}
		return rotateRight(node)
	}

	if balance < -1 {
		if getBalance(node.getRightChild()) > 0 {
			node.setRightChild(rotateRight(node.getRightChild()))
		}
		return rotateLeft(node)
	}

	return node
}

func buildLines(node *Node, prefix string, isRight bool, isOuter bool, isRoot bool, lines *[]string) {
	if node == nil {
		return
	}

	if node.getRightChild() != nil {
		spacePrefix := prefix
		newPrefix := prefix
		var newIsRight bool
		var newIsOuter bool
		if isOuter {
			if isRoot || isRight {
				newPrefix += "     "
				newIsRight = true
				newIsOuter = true
				spacePrefix += "     |"
			} else {
				newPrefix += "|    "
				newIsRight = true
				newIsOuter = false
				spacePrefix += "|    |"
			}
		} else {
			if isRoot || isRight {
				newPrefix += "     "
				newIsRight = true
				newIsOuter = false
				newIsOuter = false
				spacePrefix += "     |"
			} else {
				newPrefix += "|    "
				newIsRight = true
				newIsOuter = false
				spacePrefix += "|    |"
			}
		}
		buildLines(node.getRightChild(), newPrefix, newIsRight, newIsOuter, false, lines)
		*lines = append(*lines, spacePrefix)
	}

	line := prefix
	if !isRoot {
		line += "|-"
	} else {
		line += "  "
	}
	line += fmt.Sprintf("%3d", node.getKey())
	if node.getLeftChild() != nil || node.getRightChild() != nil {
		line += "-"
	}
	*lines = append(*lines, line)

	if node.getLeftChild() != nil {
		spacePrefix := prefix
		newPrefix := prefix
		var newIsRight bool
		var newIsOuter bool
		if isOuter {
			if isRoot || !isRight {
				newPrefix += "     "
				newIsRight = false
				newIsOuter = true
				spacePrefix += "     |"
			} else {
				newPrefix += "|    "
				newIsRight = false
				newIsOuter = false
				spacePrefix += "|    |"
			}
		} else {
			if isRoot || !isRight {
				newPrefix += "     "
				newIsRight = false
				newIsOuter = false
				spacePrefix += "     |"
			} else {
				newPrefix += "|    "
				newIsRight = false
				newIsOuter = false
				spacePrefix += "|    |"
			}
		}
		*lines = append(*lines, spacePrefix)
		buildLines(node.getLeftChild(), newPrefix, newIsRight, newIsOuter, false, lines)
	}
}

func showTree() {
	lines := []string{}
	buildLines(root, "", true, true, true, &lines)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.getLeftChild()) - getHeight(node.getRightChild())
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.getLeftChild())
	rightHeight := getHeight(node.getRightChild())
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func rotateRight(y *Node) *Node {
	x := y.getLeftChild()
	T2 := x.getRightChild()

	x.setRightChild(y)
	y.setLeftChild(T2)

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.getRightChild()
	T2 := y.getLeftChild()

	y.setLeftChild(x)
	x.setRightChild(T2)

	return y
}
