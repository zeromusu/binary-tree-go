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
	return nil
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
