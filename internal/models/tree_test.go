package models

import "testing"

func setup() {
	root = nil
}

func teardown() {
	root = nil
}

func TestAddNodeInsertRoot(t *testing.T) {
	setup()
	err := AddNode(initKey)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if root == nil || root.GetKey() != initKey {
		t.Errorf("expected root key=%d, got %v", initKey, root)
	}
	teardown()
}

func TestAddNodeInsertLeftAndRight(t *testing.T) {
	setup()
	_ = AddNode(initKey)
	_ = AddNode(leftChildKey)
	_ = AddNode(rightChildKey)

	if root.GetLeftChild() == nil || root.GetLeftChild().GetKey() != leftChildKey {
		t.Errorf("expected left child=%d, got %v", leftChildKey, root.GetLeftChild())
	}

	if root.GetRightChild() == nil || root.GetRightChild().GetKey() != rightChildKey {
		t.Errorf("expected right child=%d, got %v", rightChildKey, root.GetRightChild())
	}

	teardown()
}

func TestAddNodeDuplicate(t *testing.T) {
	setup()
	_ = AddNode(initKey)
	err := AddNode(initKey)
	if err == nil {
		t.Errorf("expected error when inserting duplicate key")
	}
}
