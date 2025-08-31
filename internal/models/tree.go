package models

import "fmt"

var root *Node

func AddNode(key int) error {
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
