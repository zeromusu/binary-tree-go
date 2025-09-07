package models

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

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

func TestShowTree(t *testing.T) {
	setup()

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	addNode(10)
	addNode(5)
	addNode(15)
	addNode(3)
	addNode(7)
	addNode(13)
	addNode(18)
	addNode(20)
	addNode(16)
	addNode(11)
	addNode(14)
	addNode(8)
	addNode(6)
	addNode(4)
	addNode(1)

	showTree()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	got := buf.String()

	want := `                               |- 20
                               |
                          |- 18-
                          |    |
                          |    |- 16
                          |
                     |- 15-
                     |    |
                     |    |    |- 14
                     |    |    |
                     |    |- 13-
                     |         |
                     |         |- 11
                     |
                   10-
                     |
                     |         |-  8
                     |         |
                     |    |-  7-
                     |    |    |
                     |    |    |-  6
                     |    |
                     |-  5-
                          |
                          |    |-  4
                          |    |
                          |-  3-
                               |
                               |-  1`

	normalize := func(s string) string {
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "\n", "")
		return s
	}

	if normalize(got) != normalize(want) {
		t.Errorf("unexpected output:\nGot:\n%s\nWant:\n%s", got, want)
	}
	teardown()
}
