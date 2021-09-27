package writer

import (
	"bytes"
	"testing"
)

// cSpell:ignore mred mgreen
var buffer = []byte("\x1b[31mred\x1b[0;32mgreen\x1b[0m\n")

//nolint:errcheck
func BenchmarkColorWriter_Write(b *testing.B) {
	w := NewWriter(&bytes.Buffer{})
	for i := 0; i < b.N; i++ {
		w.Write(buffer)
	}
}

//nolint:errcheck
func BenchmarkColorWriter_WriteColor(b *testing.B) {
	w := NewWriter(&bytes.Buffer{})
	w.SetTTY(func() bool { return true })
	for i := 0; i < b.N; i++ {
		w.Write(buffer)
	}
}

func TestColorWriter_Write(t *testing.T) {
	tests := []struct {
		name   string
		buffer string
		isTTY  bool
		want   string
	}{
		// cSpell:disable
		{
			name:   "plain text",
			buffer: "text",
			want:   "text",
		},
		{
			name:   "plain text (TTY)",
			buffer: "text",
			isTTY:  true,
			want:   "text",
		},
		{
			name:   "colored (SGR) text",
			buffer: "\x1b[31mtext\x1b[0m",
			want:   "text",
		},
		{
			name:   "colored (SGR) text (TTY)",
			buffer: "\x1b[31mtext\x1b[0m",
			isTTY:  true,
			want:   "\x1b[31mtext\x1b[0m",
		},
		{
			name:   "invalid CSI",
			buffer: "\x1b[31\x1ftext\x1b[0m",
			want:   "\x1b[31\x1ftext",
		},
		{
			name:   "invalid CSI (TTY)",
			buffer: "\x1b[31\x1ftext\x1b\x1b[0m",
			isTTY:  true,
			want:   "\x1b[31\x1ftext\x1b\x1b[0m",
		},
		{
			name:   "invalid CSI with ST",
			buffer: "\x1b[31\x1b\\text\x1b[0m",
			want:   "\x1b[31\x1b\\text",
		},
		{
			name:   "invalid CSI with ST (TTY)",
			buffer: "\x1b[31\x1b\\text\x1b\x1b[0m",
			isTTY:  true,
			want:   "\x1b[31\x1b\\text\x1b\x1b[0m",
		},
		{
			name:   "sequence start",
			buffer: "\x1b",
			want:   "\x1b",
		},
		{
			name:   "sequence start (TTY)",
			buffer: "\x1b",
			isTTY:  true,
			want:   "\x1b",
		},
		{
			name:   "CSI start",
			buffer: "\x1b[",
			want:   "\x1b[",
		},
		{
			name:   "CSI start (TTY)",
			buffer: "\x1b[",
			isTTY:  true,
			want:   "\x1b[",
		},
		{
			name:   "unterminated CSI",
			buffer: "\x1b[31",
			want:   "\x1b[31",
		},
		{
			name:   "unterminated CSI (TTY)",
			buffer: "\x1b[31",
			isTTY:  true,
			want:   "\x1b[31",
		},
		{
			name:   "cursor back (CUB) with text",
			buffer: "text\x1b[4Dreplaced",
			want:   "textreplaced",
		},
		{
			name:   "cursor back (CUB) with text (TTY)",
			buffer: "text\x1b[4Dreplaced",
			isTTY:  true,
			want:   "text\x1b[4Dreplaced",
		},
		// cSpell:enable
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			w := NewWriter(b)
			w.SetTTY(func() bool { return tt.isTTY })

			if got, err := w.Write([]byte(tt.buffer)); got != len(tt.want) || err != nil {
				t.Fatalf(
					"Write() (n, err) = (%d, %v), expected (%d, <nil>), buffer %q",
					got,
					err,
					len(tt.want),
					b.String(),
				)
			}

			if b.String() != tt.want {
				t.Fatalf("Write() = %q, expected %q", b.String(), tt.want)
			}
		})
	}
}

func TestColorWriter_Writer(t *testing.T) {
	b := &bytes.Buffer{}
	w := NewWriter(b)

	if got := w.Writer(); got != b {
		t.Fatalf("ColorWriter.Writer() = %p, expected %p", got, b)
	}
}
