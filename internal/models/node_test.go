package models

import "testing"

const initKey = 10
const changedKey = 20
const leftChildKey = initKey - 1
const rightChildKey = initKey + 1

func TestCreateNode(t *testing.T) {
	n := CreateNode(initKey, nil, nil)
	if n == nil {
		t.Fatalf("CreateNode returned nil")
	}
}

func TestDeleteNode(t *testing.T) {
	n := CreateNode(initKey, nil, nil)
	result := n.DelNode()
	if result == false {
		t.Errorf("DelNode expected true, got false")
	}
}

func TestGetKey(t *testing.T) {
	n := CreateNode(initKey, nil, nil)
	if n.GetKey() != initKey {
		t.Errorf("GetKey expected key=%d, got %d", initKey, n.GetKey())
	}
}

func TestGetLeftChild(t *testing.T) {
	left := CreateNode(leftChildKey, nil, nil)
	parent := CreateNode(initKey, left, nil)

	if parent.GetLeftChild() != left {
		t.Errorf("GetLeftChild expected left=%v, got %v", left, parent.GetLeftChild())
	}
}

func TestSetLeftChild(t *testing.T) {
	parent := CreateNode(initKey, nil, nil)
	left := CreateNode(leftChildKey, nil, nil)

	parent.SetLeftChild(left)
	if parent.GetLeftChild() != left {
		t.Errorf("SetLeftChild expected left=%v, got %v", left, parent.GetLeftChild())
	}
}

func TestGetRightChild(t *testing.T) {
	right := CreateNode(rightChildKey, nil, nil)
	parent := CreateNode(initKey, nil, right)

	if parent.GetRightChild() != right {
		t.Errorf("GetRightChild expected right=%v, got %v", right, parent.GetRightChild())
	}
}

func TestSetRightChild(t *testing.T) {
	parent := CreateNode(initKey, nil, nil)
	right := CreateNode(rightChildKey, nil, nil)

	parent.SetRightChild(right)
	if parent.GetRightChild() != right {
		t.Errorf("SetRightChild expected right=%v, got %v", right, parent.GetRightChild())
	}
}
