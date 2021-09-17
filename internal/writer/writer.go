package writer

import (
	"io"
)

type ColorWriter struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *ColorWriter {
	return &ColorWriter{
		w,
	}
}

func (w *ColorWriter) Writer() io.Writer {
	return w.writer
}

func (w *ColorWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}
