package models

import "testing"

const initKey = 10
const leftChildKey = initKey - 1
const rightChildKey = initKey + 1

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
