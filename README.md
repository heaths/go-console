# Go Console APIs

[![releases](https://img.shields.io/github/v/release/heaths/go-console.svg?logo=github)](https://github.com/heaths/go-console/releases/latest)
[![reference](https://pkg.go.dev/badge/github.com/heaths/go-console.svg)](https://pkg.go.dev/github.com/heaths/go-console)
[![ci](https://github.com/heaths/go-console/actions/workflows/ci.yml/badge.svg)](https://github.com/heaths/go-console/actions/workflows/ci.yml)

These [Go](https://golang.org) Console APIs provide a level of abstraction over virtual terminals, implement testable `io.Writer`, and support color schemes.

## Example

You can create a new console using `os` streams:

```go
package main

import (
    "fmt"

    "github.com/heaths/go-console"
)

func main() {
    con := console.System()
    fmt.Fprintln(con.Stdout(), "Hello, world!")
}
```

### Fake

You can also create a new fake console that uses `bytes.Buffer` you can access from `Stdout()` as well:

```go
package main

import (
    "fmt"

    "github.com/heaths/go-console"
)

func main() {
    fake := console.Fake()
    fmt.Fprintln(fake.Stdout(), "Hello, fake!")
    fmt.Println(fake.Stdout().String())
}
```

## License

This project is licensed under the [MIT license](LICENSE.txt).
