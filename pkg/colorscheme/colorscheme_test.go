package colorscheme

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func Benchmark_function(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = foreground(red, "red")
	}
}

func Benchmark_foregroundPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = foregroundPrintf(red, "red")
	}
}

func Benchmark_foregroundFastPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = foregroundFastPrintf(red, "red")
	}
}

// cSpell:ignore mtest
func TestColorScheme_Color(t *testing.T) {
	tests := []struct {
		fn   func(string) string
		want string
	}{
		{want: "\x1b[0;30mtest\x1b[0m", fn: Black},
		{want: "\x1b[0;31mtest\x1b[0m", fn: Red},
		{want: "\x1b[0;32mtest\x1b[0m", fn: Green},
		{want: "\x1b[0;33mtest\x1b[0m", fn: Yellow},
		{want: "\x1b[0;34mtest\x1b[0m", fn: Blue},
		{want: "\x1b[0;35mtest\x1b[0m", fn: Magenta},
		{want: "\x1b[0;36mtest\x1b[0m", fn: Cyan},
		{want: "\x1b[0;37mtest\x1b[0m", fn: White},
		{want: "\x1b[0;90mtest\x1b[0m", fn: LightBlack},
		{want: "\x1b[0;91mtest\x1b[0m", fn: LightRed},
		{want: "\x1b[0;92mtest\x1b[0m", fn: LightGreen},
		{want: "\x1b[0;93mtest\x1b[0m", fn: LightYellow},
		{want: "\x1b[0;94mtest\x1b[0m", fn: LightBlue},
		{want: "\x1b[0;95mtest\x1b[0m", fn: LightMagenta},
		{want: "\x1b[0;96mtest\x1b[0m", fn: LightCyan},
		{want: "\x1b[0;97mtest\x1b[0m", fn: LightWhite},
	}

	for _, tt := range tests {
		name := runtime.FuncForPC(reflect.ValueOf(tt.fn).Pointer()).Name()
		name = name[strings.LastIndex(name, ".")+1:]

		t.Run(name, func(t *testing.T) {
			if got := tt.fn("test"); got != tt.want {
				t.Fatalf("%s() = %q, expected %q", name, got, tt.want)
			}
		})
	}
}

// For benchmark comparisons.
func foregroundPrintf(c int, s string) string {
	return fmt.Sprintf("%s%s%dm%s%s", csi, normal, c, s, reset)
}

func foregroundFastPrintf(c int, s string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", c, s)
}
