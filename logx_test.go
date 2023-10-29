package logx

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout // Keep a copy of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	w.Close()
	os.Stdout = old
	out := <-outC
	return out
}

func TestLog(t *testing.T) {
	t.Parallel()
	expected := FGRED + BGCYAN + "Test Log" + RESET + "\n"
	output := captureStdout(func() {
		Log("Test Log", FGRED, BGCYAN)
	})

	if output != expected {
		t.Errorf("Log output: got %s, expected %s", output, expected)
	}
}
