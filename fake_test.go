package console

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithStdoutTTY(t *testing.T) {
	testSetTTY(Fake(), "Stdout", t)
}

func TestWithStderrTTY(t *testing.T) {
	testSetTTY(Fake(), "Stderr", t)
}

func TestWithStdinTTY(t *testing.T) {
	testSetTTY(Fake(), "Stdin", t)
}

func TestWithSize(t *testing.T) {
	f := Fake(WithSize(120, 40))
	if width, height, err := f.Size(); err != nil {
		t.Fatalf("Size() return error %v", err)
	} else {
		if width != 120 || height != 40 {
			t.Fatalf("Size() = %d, %d, expected 120, 40", width, height)
		}
	}
}

func TestFakeConsole_Write(t *testing.T) {
	f := Fake()
	fmt.Fprintf(f, "test")

	stdout, _, _ := f.Buffers()
	if got := stdout.String(); got != "test" {
		t.Fatalf(`Write() wrote %q, expected "test"`, got)
	}
}

func TestFakeConsole_StartProgress(t *testing.T) {
	f := Fake(
		WithStderrTTY(true),
	)

	f.StartProgress("progress")
	fmt.Fprintf(f, "test")
	f.StopProgress()

	stdout, stderr, _ := f.Buffers()
	if got := stdout.String(); got != "test" {
		t.Fatalf(`Write() wrote %q, expected "test"`, got)
	}

	// github.com/briandowns/spinner@v1.18.1 always checks os.Stdout so this will never fail.
	if stderr.Len() > 0 {
		t.Fatalf(`StartProgress() wrote progress, expected none`)
	}
}

var (
	emptyValues = []reflect.Value{}
	falseValue  = reflect.ValueOf(false)
	falseValues = []reflect.Value{falseValue}
	trueValue   = reflect.ValueOf(true)
	trueValues  = []reflect.Value{trueValue}
)

var withFuncs = map[string]func(bool) FakeOption{
	"WithStdoutTTY": WithStdoutTTY,
	"WithStderrTTY": WithStderrTTY,
	"WithStdinTTY":  WithStdinTTY,
}

func testSetTTY(f *FakeConsole, s string, t *testing.T) {
	setter := reflect.ValueOf(withFuncs["With"+s+"TTY"])
	getter := reflect.ValueOf(f).MethodByName("Is" + s + "TTY")

	// Default should evaluate handle, fail, and return false.
	if getter.Call(emptyValues)[0].Bool() {
		t.Fatalf("Is%sTTY() = true, expected fallback to false", s)
	}

	fValues := []reflect.Value{reflect.ValueOf(f)}

	setter.Call(trueValues)[0].Call(fValues)
	if !getter.Call(emptyValues)[0].Bool() {
		t.Fatalf("Is%sTTY() = false, expected true", s)
	}

	setter.Call(falseValues)[0].Call(fValues)
	if getter.Call(emptyValues)[0].Bool() {
		t.Fatalf("Is%sTTY() = true, expected false", s)
	}
}
