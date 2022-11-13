package console

import (
	"fmt"
	"strings"

	"github.com/heaths/go-console/internal/ansi"
)

func (c *con) Reset() {
	if c.IsStdoutTTY() {
		// nolint:errcheck
		c.stdout.Write([]byte(ansi.Reset))
	}
}

func (c *con) ClearLine() {
	if c.IsStdoutTTY() {
		// nolint:errcheck
		c.stdout.Write([]byte(ansi.CSI + "2K"))
	}
}

func (c *con) ClearLines(rows int) {
	if c.IsStdoutTTY() {
		// More efficient to write once.
		s := strings.Repeat(ansi.CSI+"2K"+ansi.CSI+"1A", rows)

		// nolint:errcheck
		c.stdout.Write([]byte(s))
	}
}

func (c *con) ClearScreen() {
	if c.IsStdoutTTY() {
		// nolint:errcheck
		c.stdout.Write([]byte(ansi.CSI + "2J"))
		c.MoveCursor(1, 1)
	}
}

func (c *con) StartAlternativeScreenBuffer() {
	if c.IsStdoutTTY() {
		// nolint:errcheck
		c.stdout.Write([]byte(ansi.CSI + "?1049h"))
	}
}

func (c *con) StopAlternativeScreenBuffer() {
	if c.IsStdoutTTY() {
		// nolint:errcheck
		c.stdout.Write([]byte(ansi.CSI + "?1049l"))
	}
}

func (c *con) MoveCursor(row, column int) {
	if c.IsStdoutTTY() {
		fmt.Fprintf(c.stdout, ansi.CSI+"%d;%dH", row, column)
	}
}

func (c *con) CursorUp(rows int) {
	if c.IsStdoutTTY() {
		fmt.Fprintf(c.stdout, ansi.CSI+"%dA", rows)
	}
}

func (c *con) CursorDown(rows int) {
	if c.IsStdoutTTY() {
		fmt.Fprintf(c.stdout, ansi.CSI+"%dB", rows)
	}
}

func (c *con) CursorForward(columns int) {
	if c.IsStdoutTTY() {
		fmt.Fprintf(c.stdout, ansi.CSI+"%dC", columns)
	}
}

func (c *con) CursorBack(columns int) {
	if c.IsStdoutTTY() {
		fmt.Fprintf(c.stdout, ansi.CSI+"%dD", columns)
	}
}

func (c *con) CursorColumn(column int) {
	if c.IsStdoutTTY() {
		fmt.Fprintf(c.stdout, ansi.CSI+"%dG", column)
	}
}
