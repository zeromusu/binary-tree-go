package models

import "fmt"

var root *Node

func AddNode(key int) error {
	if root == nil {
		root = CreateNode(key)
		return nil
	}
	node := root.SearchNode(key)
	if node != nil {
		return fmt.Errorf("%d is already inserted", key)
	}

	node = CreateNode(key)
	cursor := root
	for {
		if cursor.GetKey() > key {
			if cursor.GetLeftChild() == nil {
				cursor.SetLeftChild(node)
				break
			}
			cursor = cursor.GetLeftChild()
		} else {
			if cursor.GetRightChild() == nil {
				cursor.SetRightChild(node)
				break
			}
			cursor = cursor.GetRightChild()
		}
	}
	return nil
}
