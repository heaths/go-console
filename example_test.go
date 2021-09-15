package console_test

import (
	"fmt"

	"github.com/heaths/go-console"
)

func ExampleSystem() {
	con := console.System()
	fmt.Fprintln(con.Stdout(), "Hello, world!")

	// Output: Hello, world!
}

func ExampleFake() {
	fake := console.Fake()
	fmt.Fprintf(fake.Stdout(), "Hello, fake!")
	fmt.Println(fake.Stdout().String())

	// Output: Hello, fake!
}
