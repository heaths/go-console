package console

import (
	"bytes"

	"github.com/heaths/go-console/pkg/colorscheme"
)

type FakeConsole struct {
	*con
}

type FakeOption func(*FakeConsole)

func Fake(opts ...FakeOption) *FakeConsole {
	c := &con{
		stdout: &bytes.Buffer{},
		stderr: &bytes.Buffer{},
		stdin:  &bytes.Buffer{},
	}

	f := &FakeConsole{c}

	for _, opt := range opts {
		opt(f)
	}

	if c.cs == nil {
		c.cs = colorscheme.New(colorscheme.WithTTY(c.IsStdoutTTY))
	}

	return f
}

func (f *FakeConsole) Buffers() (stdout, stderr, stdin *bytes.Buffer) {
	stdout = f.stdout.(*bytes.Buffer)
	stderr = f.stderr.(*bytes.Buffer)
	stdin = f.stdin.(*bytes.Buffer)
	return
}

func (f *FakeConsole) Write(p []byte) (n int, err error) {
	return f.stdout.Write(p)
}

func WithStdout(stdout *bytes.Buffer) FakeOption {
	return func(f *FakeConsole) {
		f.stdout = stdout
	}
}

func WithStdoutTTY(tty bool) FakeOption {
	return func(f *FakeConsole) {
		f.stdoutOverride = &tty
	}
}

func WithStderr(stderr *bytes.Buffer) FakeOption {
	return func(f *FakeConsole) {
		f.stderr = stderr
	}
}

func WithStderrTTY(tty bool) FakeOption {
	return func(f *FakeConsole) {
		f.stderrOverride = &tty
	}
}

func WithStdin(stdin *bytes.Buffer) FakeOption {
	return func(f *FakeConsole) {
		f.stdin = stdin
	}
}

func WithStdinTTY(tty bool) FakeOption {
	return func(f *FakeConsole) {
		f.stdinOverride = &tty
	}
}

func WithSize(width, height int) FakeOption {
	if width < 0 {
		panic("width cannot be less than 0")
	}
	if height < 0 {
		panic("height cannot be less than 0")
	}
	return func(f *FakeConsole) {
		f.sizeOverride = &struct {
			Width  int
			Height int
		}{
			Width:  width,
			Height: height,
		}
	}
}

func WithColorScheme(cs *colorscheme.ColorScheme) FakeOption {
	return func(f *FakeConsole) {
		f.cs = cs
	}
}
