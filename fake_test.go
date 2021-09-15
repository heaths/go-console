package console

import (
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

var (
	emptyValues = []reflect.Value{}
	falseValue  = reflect.ValueOf(false)
	falseValues = []reflect.Value{falseValue}
	trueValue   = reflect.ValueOf(true)
	trueValues  = []reflect.Value{trueValue}
)

func testSetTTY(f *FakeConsole, s string, t *testing.T) {
	setter := reflect.ValueOf(f).MethodByName("Set" + s + "TTY")
	getter := reflect.ValueOf(f).MethodByName("Is" + s + "TTY")

	setter.Call(trueValues)
	if !getter.Call(emptyValues)[0].Bool() {
		t.Fatalf("Is%sTTY() = false, expected true", s)
	}

	setter.Call(falseValues)
	if getter.Call(emptyValues)[0].Bool() {
		t.Fatalf("Is%sTTY() = true, expected false", s)
	}
}
