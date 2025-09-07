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
	err := addNode(initKey)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if root == nil || root.getKey() != initKey {
		t.Errorf("expected root key=%d, got %v", initKey, root)
	}
	teardown()
}

func TestAddNodeInsertLeftAndRight(t *testing.T) {
	setup()
	_ = addNode(initKey)
	_ = addNode(leftChildKey)
	_ = addNode(rightChildKey)

	if root.getLeftChild() == nil || root.getLeftChild().getKey() != leftChildKey {
		t.Errorf("expected left child=%d, got %v", leftChildKey, root.getLeftChild())
	}

	if root.getRightChild() == nil || root.getRightChild().getKey() != rightChildKey {
		t.Errorf("expected right child=%d, got %v", rightChildKey, root.getRightChild())
	}

	teardown()
}

func TestAddNodeDuplicate(t *testing.T) {
	setup()
	_ = addNode(initKey)
	err := addNode(initKey)
	if err == nil {
		t.Errorf("expected error when inserting duplicate key")
	}
}
