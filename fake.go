package console

import (
	"bytes"
)

type FakeConsole struct {
	Console
}

func Fake() *FakeConsole {
	return &FakeConsole{
		Console{
			stdout: &bytes.Buffer{},
			stderr: &bytes.Buffer{},
			stdin:  &bytes.Buffer{},
		},
	}
}

func (f *FakeConsole) Stdout() *bytes.Buffer {
	return f.stdout.(*bytes.Buffer)
}

func (f *FakeConsole) SetStdoutTTY(tty bool) {
	f.stdoutOverride = &tty
}

func (f *FakeConsole) Stderr() *bytes.Buffer {
	return f.stderr.(*bytes.Buffer)
}

func (f *FakeConsole) SetStderrTTY(tty bool) {
	f.stderrOverride = &tty
}

func (f *FakeConsole) Stdin() *bytes.Buffer {
	return f.stdin.(*bytes.Buffer)
}

func (f *FakeConsole) SetStdinTTY(tty bool) {
	f.stdinOverride = &tty
}

func (f *FakeConsole) Write(p []byte) (n int, err error) {
	return f.stdout.Write(p)
}
