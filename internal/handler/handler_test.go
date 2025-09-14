package handler

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCLI(t *testing.T) {
	input := "insert 5\ninsert 5\nget 5\nget 3\ndelete 3\ndelete 5\nshow\nexit\n"

	oldStdin := os.Stdin
	rIn, wIn, _ := os.Pipe()
	os.Stdin = rIn

	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	go func() {
		defer wIn.Close()
		wIn.Write([]byte(input))
	}()

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		defer close(done)
		io.Copy(&buf, rOut)
	}()

	RunCLI()

	os.Stdin = oldStdin
	wOut.Close()
	os.Stdout = oldStdout
	<-done

	got := buf.String()

	wantContains := []string{
		"inserted 5",
		"5 is already exists",
		"5 found",
		"3 not found",
		"3 not found",
		"5 deleted",
	}

	for _, str := range wantContains {
		if !bytes.Contains([]byte(got), []byte(str)) {
			t.Errorf("expected output to contain %q, got:\n%s", str, got)
		}
	}
}
