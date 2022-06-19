package console_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/heaths/go-console"
	"github.com/heaths/go-console/pkg/colorscheme"
)

func Example() {
	stdin := bytes.NewBufferString("31\tred\n32\tgreen\n")

	// Set up fake console with stdin, and stdout as TTY.
	fake := console.Fake(
		console.WithStdin(stdin),
		console.WithStdoutTTY(true),
	)

	// Scan color codes and descriptions from fake stdin.
	scanner := bufio.NewScanner(fake.Stdin())
	for scanner.Scan() {
		var color int
		var desc string

		// Write scanned color codes to fake stdout.
		if _, err := fmt.Sscanf(scanner.Text(), "%d %s", &color, &desc); err == nil {
			fmt.Fprintf(fake.Stdout(), "\x1b[0;%dm%s\x1b[0m", color, desc)
		}
	}

	// Doubly escape fake stdout and write to real stdout to assert output.
	stdout, _, _ := fake.Buffers()
	s := strings.ReplaceAll(stdout.String(), "\x1b", `\x1b`)
	fmt.Println(s)

	// Output: \x1b[0;31mred\x1b[0m\x1b[0;32mgreen\x1b[0m
}

func ExampleSystem() {
	// Create console from system streams.
	con := console.System()
	fmt.Fprintln(con.Stdout(), "Hello, world!")

	// Output: Hello, world!
}

func ExampleFake() {
	// Create fake console from buffers.
	fake := console.Fake()
	fmt.Fprintf(fake.Stdout(), "Hello, fake!")

	stdout, _, _ := fake.Buffers()
	fmt.Println(stdout.String())

	// Output: Hello, fake!
}

func ExampleConsole_ColorScheme() {
	// Set up a console with stdout redirected, but not stderr.
	fake := console.Fake(
		console.WithStdoutTTY(false),
		console.WithStderrTTY(true),
	)

	cs := fake.ColorScheme()

	// Console.ColorScheme affects the Console i.e., stdout, so this isn't colored.
	fmt.Fprintf(fake.Stderr(), "%s\n", cs.Red("Error: not red!"))

	// Clone the ColorScheme to check if stderr is redirected or not.
	cs = fake.ColorScheme().Clone(
		colorscheme.WithTTY(fake.IsStderrTTY),
	)
	fmt.Fprintf(fake.Stderr(), "%s\n", cs.Red("Error: red alert!"))

	// Doubly escape fake stdout and write to real stdout to assert output.
	_, stderr, _ := fake.Buffers()
	s := strings.ReplaceAll(stderr.String(), "\x1b", `\x1b`)
	fmt.Println(s)

	// Output:
	// Error: not red!
	// \x1b[0;31mError: red alert!\x1b[0m
}
