package console

import (
	"io"
	"os"

	"github.com/heaths/go-console/internal/writer"
	"github.com/mattn/go-isatty"
)

type Console struct {
	stdout *writer.ColorWriter
	stderr *writer.ColorWriter
	stdin  io.Reader

	stdoutOverride *bool
	stderrOverride *bool
	stdinOverride  *bool
}

func System() *Console {
	return newConsole(
		os.Stdout,
		os.Stderr,
		os.Stdin,
	)
}

func (c *Console) Stdout() io.Writer {
	return c.stdout
}

func (c *Console) IsStdoutTTY() bool {
	if c.stdoutOverride != nil {
		return *c.stdoutOverride
	}

	if w, ok := c.stdout.Writer().(*os.File); ok {
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

	if w, ok := c.stderr.Writer().(*os.File); ok {
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

func (c *Console) Write(p []byte) (n int, err error) {
	return c.stdout.Write(p)
}

func newConsole(stdout, stderr io.Writer, stdin io.Reader) *Console {
	// Set up writer to remove CSI sequences.
	// TODO: Consider a way to avoid writing CSI sequences if target Writer is not a TTY.
	c := &Console{
		stdout: writer.NewWriter(stdout),
		stderr: writer.NewWriter(stderr),
		stdin:  stdin,
	}

	c.stdout.SetTTY(c.IsStdoutTTY)
	c.stderr.SetTTY(c.IsStderrTTY)

	return c
}
