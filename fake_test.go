package console

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetStdoutTTY(t *testing.T) {
	testSetTTY(Fake(), "Stdout", t)
}

func TestSetStderrTTY(t *testing.T) {
	testSetTTY(Fake(), "Stderr", t)
}

func TestSetStdinTTY(t *testing.T) {
	testSetTTY(Fake(), "Stdin", t)
}

func TestWrite(t *testing.T) {
	f := Fake()
	fmt.Fprintf(f, "test")
	if got := f.Stdout().String(); got != "test" {
		t.Fatalf(`Write() wrote %q, expected "test"`, got)
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
