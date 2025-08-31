package models

import "testing"

const initKey = 10
const changedKey = 20
const leftChildKey = initKey - 1
const rightChildKey = initKey + 1

func TestCreateNode(t *testing.T) {
	n := CreateNode(initKey, nil, nil, nil)
	if n == nil {
		t.Fatalf("CreateNode returned nil")
	}
}

func TestDeleteNode(t *testing.T) {
	n := CreateNode(initKey, nil, nil, nil)
	result := n.DelNode()
	if result == false {
		t.Errorf("DelNode expected true, got false")
	}
}

func TestGetKey(t *testing.T) {
	n := CreateNode(initKey, nil, nil, nil)
	if n.GetKey() != initKey {
		t.Errorf("GetKey expected key=%d, got %d", initKey, n.GetKey())
	}
}

func TestSetKey(t *testing.T) {
	n := CreateNode(initKey, nil, nil, nil)
	n.SetKey(changedKey)
	if n.GetKey() != changedKey {
		t.Errorf("SetKey expected key=%d, got %d", changedKey, n.GetKey())
	}
}

func TestGetParent(t *testing.T) {
	parent := CreateNode(initKey, nil, nil, nil)
	child := CreateNode(leftChildKey, parent, nil, nil)

	if child.GetParent() != parent {
		t.Errorf("GetParent expected parent=%v, got %v", parent, child.GetParent())
	}
}

func TestSetParent(t *testing.T) {
	parent := CreateNode(initKey, nil, nil, nil)
	child := CreateNode(leftChildKey, nil, nil, nil)

	child.SetParent(parent)
	if child.GetParent() != parent {
		t.Errorf("SetParent expected parent=%v, got %v", parent, child.GetParent())
	}
}

func TestGetLeftChild(t *testing.T) {
	left := CreateNode(leftChildKey, nil, nil, nil)
	parent := CreateNode(initKey, nil, left, nil)

	if parent.GetLeftChild() != left {
		t.Errorf("GetLeftChild expected left=%v, got %v", left, parent.GetLeftChild())
	}
}

func TestSetLeftChild(t *testing.T) {
	parent := CreateNode(initKey, nil, nil, nil)
	left := CreateNode(leftChildKey, parent, nil, nil)

	parent.SetLeftChild(left)
	if parent.GetLeftChild() != left {
		t.Errorf("SetLeftChild expected left=%v, got %v", left, parent.GetLeftChild())
	}
}

func TestGetRightChild(t *testing.T) {
	right := CreateNode(rightChildKey, nil, nil, nil)
	parent := CreateNode(initKey, nil, nil, right)

	if parent.GetRightChild() != right {
		t.Errorf("GetRightChild expected right=%v, got %v", right, parent.GetRightChild())
	}
}

func TestSetRightChild(t *testing.T) {
	parent := CreateNode(initKey, nil, nil, nil)
	right := CreateNode(rightChildKey, parent, nil, nil)

	parent.SetRightChild(right)
	if parent.GetRightChild() != right {
		t.Errorf("SetRightChild expected right=%v, got %v", right, parent.GetRightChild())
	}
}
