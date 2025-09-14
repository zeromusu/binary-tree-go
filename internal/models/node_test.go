package models

import "testing"

const initKey = 10
const changedKey = initKey + 10
const leftChildKey = initKey - 5
const rightChildKey = initKey + 5
const noExistingKey = -1

func TestCreateNode(t *testing.T) {
	n := createNode(initKey)
	if n == nil {
		t.Fatalf("CreateNode returned nil")
	}
}

func TestDeleteNode(t *testing.T) {
	n := createNode(initKey)
	result := n.delNode()
	if result == false {
		t.Errorf("DelNode expected true, got false")
	}
}

func TestSearchNode(t *testing.T) {
	n := createNode(initKey)
	searchedNode := n.searchNode(noExistingKey)
	if searchedNode != nil {
		t.Errorf("SearchNode for no existing key expected nil, got %v", searchedNode)
	}

	left := createNode(leftChildKey)
	n.setLeftChild(left)
	leftLeft := createNode(leftChildKey - 1)
	leftRight := createNode(leftChildKey + 1)
	left.setLeftChild(leftLeft)
	left.setRightChild(leftRight)
	searchedNode = n.searchNode(leftChildKey - 1)
	if searchedNode != leftLeft {
		t.Errorf("SearchNode for %d expected %v, got %v", leftChildKey-1, leftLeft, searchedNode)
	}
	searchedNode = n.searchNode(leftChildKey + 1)
	if searchedNode != leftRight {
		t.Errorf("SearchNode for %d expected %v, got %v", leftChildKey+1, leftRight, searchedNode)
	}
}

func TestSetAndGetKey(t *testing.T) {
	n := createNode(initKey)
	if n.getKey() != initKey {
		t.Errorf("GetKey expected key=%d, got %d", initKey, n.getKey())
	}
	n.setKey(changedKey)
	if n.getKey() != changedKey {
		t.Errorf("SetKey expected key=%d, got %d", changedKey, n.getKey())
	}
}

func TestSetAndGetLeftChild(t *testing.T) {
	parent := createNode(initKey)
	left := createNode(leftChildKey)

	parent.setLeftChild(left)
	if parent.getLeftChild() != left {
		t.Errorf("SetLeftChild expected left=%v, got %v", left, parent.getLeftChild())
	}
}

func TestSetAndGetRightChild(t *testing.T) {
	parent := createNode(initKey)
	right := createNode(rightChildKey)

	parent.setRightChild(right)
	if parent.getRightChild() != right {
		t.Errorf("SetRightChild expected right=%v, got %v", right, parent.getRightChild())
	}
}
