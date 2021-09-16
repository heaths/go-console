package console

import (
	"bytes"
)

type FakeConsole struct {
	Console
}

type FakeOption func(*FakeConsole)

func Fake(opts ...FakeOption) *FakeConsole {
	f := &FakeConsole{
		Console{
			stdout: &bytes.Buffer{},
			stderr: &bytes.Buffer{},
			stdin:  &bytes.Buffer{},
		},
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *FakeConsole) Stdout() *bytes.Buffer {
	return f.stdout.(*bytes.Buffer)
}

func (f *FakeConsole) Stderr() *bytes.Buffer {
	return f.stderr.(*bytes.Buffer)
}

func (f *FakeConsole) Stdin() *bytes.Buffer {
	return f.stdin.(*bytes.Buffer)
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
