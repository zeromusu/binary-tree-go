package models

import "testing"

const initKey = 10
const leftChildKey = initKey - 5
const rightChildKey = initKey + 5
const noExistingKey = -1

func TestCreateNode(t *testing.T) {
	n := CreateNode(initKey)
	if n == nil {
		t.Fatalf("CreateNode returned nil")
	}
}

func TestDeleteNode(t *testing.T) {
	n := CreateNode(initKey)
	result := n.DelNode()
	if result == false {
		t.Errorf("DelNode expected true, got false")
	}
}

func TestSearchNode(t *testing.T) {
	n := CreateNode(initKey)
	searchedNode := n.SearchNode(noExistingKey)
	if searchedNode != nil {
		t.Errorf("SearchNode for no existing key expected nil, got %v", searchedNode)
	}

	left := CreateNode(leftChildKey)
	n.SetLeftChild(left)
	leftLeft := CreateNode(leftChildKey - 1)
	leftRight := CreateNode(leftChildKey + 1)
	left.SetLeftChild(leftLeft)
	left.SetRightChild(leftRight)
	searchedNode = n.SearchNode(leftChildKey - 1)
	if searchedNode != leftLeft {
		t.Errorf("SearchNode for %d expected %v, got %v", leftChildKey-1, leftLeft, searchedNode)
	}
	searchedNode = n.SearchNode(leftChildKey + 1)
	if searchedNode != leftRight {
		t.Errorf("SearchNode for %d expected %v, got %v", leftChildKey+1, leftRight, searchedNode)
	}
}

func TestGetKey(t *testing.T) {
	n := CreateNode(initKey)
	if n.GetKey() != initKey {
		t.Errorf("GetKey expected key=%d, got %d", initKey, n.GetKey())
	}
}

func TestSetAndGetLeftChild(t *testing.T) {
	parent := CreateNode(initKey)
	left := CreateNode(leftChildKey)

	parent.SetLeftChild(left)
	if parent.GetLeftChild() != left {
		t.Errorf("SetLeftChild expected left=%v, got %v", left, parent.GetLeftChild())
	}
}

func TestSetAndGetRightChild(t *testing.T) {
	parent := CreateNode(initKey)
	right := CreateNode(rightChildKey)

	parent.SetRightChild(right)
	if parent.GetRightChild() != right {
		t.Errorf("SetRightChild expected right=%v, got %v", right, parent.GetRightChild())
	}
}
