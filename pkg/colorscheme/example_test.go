package colorscheme_test

import (
	"fmt"
	"strings"

	"github.com/heaths/go-console"
	"github.com/heaths/go-console/pkg/colorscheme"
)

func ExampleColorScheme() {
	fake := console.Fake()

	cs := colorscheme.New(
		colorscheme.WithTTY(fake.IsStdoutTTY),
	)

	greeting := cs.ColorFunc("green")
	fmt.Fprintf(fake.Stdout(), "%s", greeting("Hello, world!"))
	fmt.Println(fake.Stdout().String())

	// Output: Hello, world!
}

func ExampleColorScheme_color() {
	fake := console.Fake(
		console.WithStdoutTTY(true),
	)

	cs := colorscheme.New(
		colorscheme.WithTTY(fake.IsStdoutTTY),
	)

	greeting := cs.ColorFunc("green")
	fmt.Fprintf(fake.Stdout(), "%s", greeting("Hello, world!"))

	// Doubly escape fake stdout and write to real stdout to assert output.
	s := strings.ReplaceAll(fake.Stdout().String(), "\x1b", `\x1b`)
	fmt.Println(s)

	// Output: \x1b[0;32mHello, world!\x1b[0m
}
