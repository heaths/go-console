package console

import (
	"bytes"

	"github.com/heaths/go-console/internal/writer"
)

type FakeConsole struct {
	Console
}

type FakeOption func(*FakeConsole)

func Fake(opts ...FakeOption) *FakeConsole {
	f := &FakeConsole{
		Console{
			stdout: writer.NewWriter(&bytes.Buffer{}),
			stderr: writer.NewWriter(&bytes.Buffer{}),
			stdin:  &bytes.Buffer{},
		},
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *FakeConsole) Stdout() *bytes.Buffer {
	return f.stdout.Writer().(*bytes.Buffer)
}

func (f *FakeConsole) Stderr() *bytes.Buffer {
	return f.stderr.Writer().(*bytes.Buffer)
}

func (f *FakeConsole) Stdin() *bytes.Buffer {
	return f.stdin.(*bytes.Buffer)
}

func (f *FakeConsole) Write(p []byte) (n int, err error) {
	return f.stdout.Write(p)
}

func WithStdout(stdout *bytes.Buffer) FakeOption {
	return func(f *FakeConsole) {
		f.stdout = writer.NewWriter(stdout)
	}
}

func WithStdoutTTY(tty bool) FakeOption {
	return func(f *FakeConsole) {
		f.stdoutOverride = &tty
	}
}

func WithStderr(stderr *bytes.Buffer) FakeOption {
	return func(f *FakeConsole) {
		f.stderr = writer.NewWriter(stderr)
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
