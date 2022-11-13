package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/heaths/go-console"
)

func main() {
	con := console.System()
	cs := con.ColorScheme()

	con.StartAlternativeScreenBuffer()
	defer con.StopAlternativeScreenBuffer()

	con.MoveCursor(2, 2)
	fmt.Fprintln(con, "Shall we play a game?")

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	for {
		con.ClearLine()
		con.CursorColumn(2)
		fmt.Fprintf(con, cs.LightBlack("Launching in %d..."), int(timeout.Seconds()))

		select {
		case <-time.After(time.Second):
			timeout -= time.Second
		case <-ctx.Done():
			con.StopAlternativeScreenBuffer()
			cancel()

			return
		}
	}
}
