package writer

import (
	"io"
)

type ColorWriter struct {
	writer io.Writer
	isTTY  func() bool
}

func New(w io.Writer) *ColorWriter {
	return &ColorWriter{
		writer: w,
		isTTY:  isTTY,
	}
}

func (w *ColorWriter) SetTTY(isTTY func() bool) {
	w.isTTY = isTTY
}

func (w *ColorWriter) Writer() io.Writer {
	return w.writer
}

func (w *ColorWriter) Write(p []byte) (n int, err error) {
	// If writer is TTY, just pass through.
	if w.isTTY() {
		return w.writer.Write(p)
	}

	// Scan for and remove CSI sequences.
	const (
		ESC = '\x1b'
		CSI = '['
		SGR = 'm'
	)

	var (
		bytesWritten int

		rangeStart = 0
		seqStart   = -1
		isCSI      = false
	)

	for i, b := range p {
		switch {
		// Start of escape sequence.
		case b == ESC:
			seqStart = i
			isCSI = false

		// Start of CSI sequence.
		case b == CSI && seqStart == i-1:
			isCSI = true

		// CSI parameter bytes.
		case b >= 0x30 && b <= 0x3f && isCSI:

		// CSI intermediate bytes.
		case b >= 0x20 && b <= 0x2f && isCSI:

		// CSI terminal byte: SGR.
		case b == SGR && isCSI:
			// TODO: Consider whether to only ignore SGR or all CSI terminators.
			fallthrough

		// CSI terminal byte.
		case b >= 0x40 && b <= 0x7e && isCSI:
			// Write buffer up til start of escape sequence.
			if bytesWritten, err = w.writer.Write(p[rangeStart:seqStart]); err != nil {
				return
			}

			n += bytesWritten
			rangeStart = i + 1

			fallthrough

		// Reset.
		default:
			seqStart = -1
			isCSI = false
		}
	}

	// Write remaining buffer.
	if bytesWritten, err = w.writer.Write(p[rangeStart:]); err != nil {
		return
	}

	n += bytesWritten
	return
}

func isTTY() bool {
	return false
}
