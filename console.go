package console

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/heaths/go-console/pkg/colorscheme"
	"github.com/mattn/go-isatty"
)

type Console interface {
	Stdout() io.Writer
	Stderr() io.Writer
	Stdin() io.Reader
	IsStdoutTTY() bool
	IsStderrTTY() bool
	IsStdinTTY() bool

	io.Writer

	ColorScheme() *colorscheme.ColorScheme
	Reset()

	StartProgress(label string, opts ...ProgressOption)
	StopProgress()

	ClearLine()
	ClearLines(rows int)
	ClearScreen()
	StartAlternativeScreenBuffer()
	StopAlternativeScreenBuffer()

	MoveCursor(rows, columns int)
	CursorUp(rows int)
	CursorDown(rows int)
	CursorForward(columns int)
	CursorBack(columns int)
	CursorColumn(column int)
}

type con struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader

	stdoutOverride *bool
	stderrOverride *bool
	stdinOverride  *bool

	cs *colorscheme.ColorScheme

	progress        *spinner.Spinner
	progressEnabled bool
	progressLock    sync.Mutex
	progressMin     <-chan time.Time
}

func System() Console {
	c := &con{
		stdout: os.Stdout,
		stderr: os.Stderr,
		stdin:  os.Stdin,

		progressEnabled: true,
	}

	c.cs = colorscheme.New(colorscheme.WithTTY(c.IsStdoutTTY))

	return c
}

func (c *con) Stdout() io.Writer {
	return c.stdout
}

func (c *con) IsStdoutTTY() bool {
	if c.stdoutOverride != nil {
		return *c.stdoutOverride
	}

	if w, ok := c.stdout.(*os.File); ok {
		return isatty.IsTerminal(w.Fd())
	}

	return false
}

func (c *con) Stderr() io.Writer {
	return c.stderr
}

func (c *con) IsStderrTTY() bool {
	if c.stderrOverride != nil {
		return *c.stderrOverride
	}

	if w, ok := c.stderr.(*os.File); ok {
		return isatty.IsTerminal(w.Fd())
	}

	return false
}

func (c *con) Stdin() io.Reader {
	return c.stdin
}

func (c *con) IsStdinTTY() bool {
	if c.stdinOverride != nil {
		return *c.stdinOverride
	}

	if w, ok := c.stdin.(*os.File); ok {
		return isatty.IsTerminal(w.Fd())
	}

	return false
}

// Write implements Writer on the console and calls Write on Stdout.
func (c *con) Write(p []byte) (n int, err error) {
	return c.stdout.Write(p)
}

// ColorScheme gets the color scheme for the console i.e., Stdout.
func (c *con) ColorScheme() *colorscheme.ColorScheme {
	return c.cs
}
