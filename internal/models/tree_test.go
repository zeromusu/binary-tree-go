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
	err := AddNode(initKey)
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
	_ = AddNode(initKey)
	_ = AddNode(leftChildKey)
	_ = AddNode(rightChildKey)

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
	_ = AddNode(initKey)
	err := AddNode(initKey)
	if err == nil {
		t.Errorf("expected error when inserting duplicate key")
	}
}

func TestGetNode(t *testing.T) {
	setup()

	AddNode(5)
	AddNode(3)
	AddNode(8)
	AddNode(1)
	AddNode(4)
	AddNode(6)
	AddNode(10)

	tests := []struct {
		key  int
		want bool
	}{
		{5, true},
		{3, true},
		{10, true},
		{7, false},
		{0, false},
	}

	for _, tt := range tests {
		got := FindNode(tt.key)
		if got != tt.want {
			t.Errorf("FindNode(%d) = %v, want %v", tt.key, got, tt.want)
		}
	}
}

func TestDeleteNode(t *testing.T) {
	setup()

	AddNode(10)
	AddNode(5)
	AddNode(15)
	AddNode(3)
	AddNode(7)
	AddNode(13)
	AddNode(18)

	gotBefore := captureOutput(func() {
		ShowTree()
	})

	deleteKey := 15
	if err := DeleteNode(deleteKey); err != nil {
		t.Fatalf("deleteNode failed: %v", err)
	}

	gotAfter := captureOutput(func() {
		ShowTree()
	})

	want := `
       |- 18-
       |    |
       |    |- 13
       |
     10-
       |
       |    |- 7
       |    |
       |- 5-
            |
            |- 3
`

	normalize := func(s string) string {
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "\n", "")
		return s
	}

	if normalize(gotAfter) != normalize(want) {
		t.Errorf("unexpected tree after deleteNode(%d):\nGot:\n%s\nWant:\n%s", deleteKey, gotAfter, want)
	}

	if gotBefore == gotAfter {
		t.Errorf("tree did not change after deleteNode(%d)", deleteKey)
	}
}

func TestRebalanceRotations(t *testing.T) {
	setup()

	AddNode(1)
	AddNode(2)
	AddNode(3)

	if root == nil || root.getKey() != 2 {
		t.Fatalf("RR rotation failed: root want=2, got=%v", root.getKey())
	}
	if root.getLeftChild() == nil || root.getLeftChild().getKey() != 1 {
		t.Fatalf("RR rotation left child wrong")
	}
	if root.getRightChild() == nil || root.getRightChild().getKey() != 3 {
		t.Fatalf("RR rotation right child wrong")
	}

	setup()

	AddNode(3)
	AddNode(2)
	AddNode(1)

	if root == nil || root.getKey() != 2 {
		t.Fatalf("LL rotation failed: root want=2, got=%v", root.getKey())
	}
	if root.getLeftChild() == nil || root.getLeftChild().getKey() != 1 {
		t.Fatalf("LL rotation left child wrong")
	}
	if root.getRightChild() == nil || root.getRightChild().getKey() != 3 {
		t.Fatalf("LL rotation right child wrong")
	}

	setup()

	AddNode(3)
	AddNode(1)
	AddNode(2)

	if root == nil || root.getKey() != 2 {
		t.Fatalf("LR rotation failed: root want=2, got=%v", root.getKey())
	}

	setup()

	AddNode(1)
	AddNode(3)
	AddNode(2)

	if root == nil || root.getKey() != 2 {
		t.Fatalf("RL rotation failed: root want=2, got=%v", root.getKey())
	}
}

func TestRebalance(t *testing.T) {
	setup()

	AddNode(1)
	AddNode(2)
	AddNode(3)
	AddNode(4)
	AddNode(5)

	gotAfter := captureOutput(func() {
		ShowTree()
	})

	want := `
       |- 5
       |
  |- 4-
  |    |
  |    |- 3
  |
 2-
  |
  |- 1
`

	normalize := func(s string) string {
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "\n", "")
		return s
	}

	if normalize(gotAfter) != normalize(want) {
		t.Errorf("unexpected output:\nGot:\n%s\nWant:\n%s", gotAfter, want)
	}

	teardown()
}

func TestShowTree(t *testing.T) {
	setup()

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	AddNode(10)
	AddNode(5)
	AddNode(15)
	AddNode(3)
	AddNode(7)
	AddNode(13)
	AddNode(18)
	AddNode(20)
	AddNode(16)
	AddNode(11)
	AddNode(14)
	AddNode(8)
	AddNode(6)
	AddNode(4)
	AddNode(1)

	ShowTree()

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

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
