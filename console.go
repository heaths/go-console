package console

import (
	"io"
	"os"

	"github.com/heaths/go-console/pkg/colorscheme"
	"github.com/mattn/go-isatty"
)

type Console struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader

	stdoutOverride *bool
	stderrOverride *bool
	stdinOverride  *bool

	cs *colorscheme.ColorScheme
}

func System() *Console {
	c := &Console{
		stdout: os.Stdout,
		stderr: os.Stderr,
		stdin:  os.Stdin,
	}

	c.cs = colorscheme.New(colorscheme.WithTTY(c.IsStdoutTTY))

	return c
}

func (c *Console) Stdout() io.Writer {
	return c.stdout
}

func (c *Console) IsStdoutTTY() bool {
	if c.stdoutOverride != nil {
		return *c.stdoutOverride
	}

	if w, ok := c.stdout.(*os.File); ok {
		return isatty.IsTerminal(w.Fd())
	}

	return false
}

func (c *Console) Stderr() io.Writer {
	return c.stderr
}

func (c *Console) IsStderrTTY() bool {
	if c.stderrOverride != nil {
		return *c.stderrOverride
	}

	if w, ok := c.stderr.(*os.File); ok {
		return isatty.IsTerminal(w.Fd())
	}

	return false
}

func (c *Console) Stdin() io.Reader {
	return c.stdin
}

func (c *Console) IsStdinTTY() bool {
	if c.stdinOverride != nil {
		return *c.stdinOverride
	}

	if w, ok := c.stdin.(*os.File); ok {
		return isatty.IsTerminal(w.Fd())
	}

	return false
}

// Write implements Writer on the console and calls Write on Stdout.
func (c *Console) Write(p []byte) (n int, err error) {
	return c.stdout.Write(p)
}

// ColorScheme gets the color scheme for the console i.e., Stdout.
func (c *Console) ColorScheme() *colorscheme.ColorScheme {
	return c.cs
}
