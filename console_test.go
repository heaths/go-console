package console

import (
	"os"
	"reflect"
	"testing"
)

func TestConsole_IsStdoutTTY(t *testing.T) {
	testIsTTY("Stdout", t)
}

func TestConsole_IsStderrTTY(t *testing.T) {
	testIsTTY("Stderr", t)
}

func TestConsole_IsStdinTTY(t *testing.T) {
	testIsTTY("Stdin", t)
}

func testIsTTY(s string, t *testing.T) {
	var f *os.File
	var err error

	if f, err = os.CreateTemp("", "test"); err != nil {
		t.Fatalf("CreateTemp() error = %v", err)
	}

	defer f.Close()

	console := &con{
		stdout: f,
		stderr: f,
		stdin:  f,
	}

	getter := reflect.ValueOf(console).MethodByName("Is" + s + "TTY")

	// Default should evaluate handle, fail, and return false.
	if getter.Call(emptyValues)[0].Bool() {
		t.Fatalf("Is%sTTY() = true, expected false", s)
	}
}
