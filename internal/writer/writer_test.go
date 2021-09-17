package writer

import (
	"bytes"
	"testing"
)

func TestColorWriter_Writer(t *testing.T) {
	b := &bytes.Buffer{}
	w := NewWriter(b)

	if got := w.Writer(); got != b {
		t.Fatalf("ColorWriter.Writer() = %p, expected %p", got, b)
	}
}
