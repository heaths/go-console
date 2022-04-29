package colorscheme

import (
	"bytes"
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

func TestNew(t *testing.T) {
	got := New()
	if len(got.colors) != 0 {
		t.Fatalf("New() len(colors) = %d, expected 0", len(got.colors))
	}
}

// cSpell:ignore magain mtest
func TestColorScheme_ColorFunc(t *testing.T) {
	cs := New(WithTTY(alwaysTTY))
	fn := cs.ColorFunc("red+h")

	want := "\x1b[0;91mtest\x1b[0m"
	if got := fn("test"); got != want {
		t.Fatalf("ColorFunc()() = %q, expected %q", got, want)
	}

	if len(cs.colors) != 1 {
		t.Fatalf("len(ColorScheme.colors) = %d, expected 1", len(cs.colors))
	}
}

func TestColorScheme_ColorFunc_empty(t *testing.T) {
	cs := New()
	fn := cs.ColorFunc("")

	want := "test"
	if got := fn("test"); got != want {
		t.Fatalf("ColorFunc()() = %q, expected %q", got, want)
	}

	if len(cs.colors) != 0 {
		t.Fatalf("len(ColorScheme.colors) = %d, expected 0", len(cs.colors))
	}
}

// TestColorScheme_ColorFunc_twice tests for heaths/go-console#9.
func TestColorScheme_ColorFunc_twice(t *testing.T) {
	cs := New(WithTTY(alwaysTTY))
	fn := cs.ColorFunc("red+h")

	want := "\x1b[0;91mtest\x1b[0m\x1b[0;91magain\x1b[0m"
	if got := fn("test") + fn("again"); got != want {
		t.Fatalf("ColorFunc()() = %q, expected %q", got, want)
	}

	if len(cs.colors) != 1 {
		t.Fatalf("len(ColorScheme.colors) = %d, expected 1", len(cs.colors))
	}
}

// cSpell:ignore mtest
func TestColor(t *testing.T) {
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

func TestColorCode(t *testing.T) {
	tests := []struct {
		name  string
		style string
		want  string
	}{
		{
			name:  "empty",
			style: "",
			want:  "",
		},
		{
			name:  "off",
			style: "off",
			want:  "",
		},
		{
			name:  "reset",
			style: "reset",
			want:  "\x1b[0m",
		},
		{
			name:  "only foreground",
			style: "red+h",
			want:  "\x1b[0;91m",
		},
		{
			name:  "only background",
			style: ":green+d",
			want:  "\x1b[0;2;42m",
		},
		{
			name:  "foreground and background",
			style: "red+h:green+d",
			want:  "\x1b[0;91;2;42m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := colorCode(tt.style)
			if got := buf.String(); got != tt.want {
				t.Fatalf("colorCode() = %q, expected %q", got, tt.want)
			}
		})
	}
}

func TestColorPartCode(t *testing.T) {
	tests := []struct {
		name string
		part string
		base int
		want string
	}{
		{
			name: "empty",
			part: "",
			base: normalFG,
			want: "",
		},
		{
			name: "empty color",
			part: "+h",
			base: normalFG,
			want: "",
		},
		{
			name: "invalid",
			part: "invalid",
			base: normalFG,
			want: "0",
		},
		{
			name: "named color (foreground)",
			part: "red",
			base: normalFG,
			want: "31",
		},
		{
			name: "light named color (foreground)",
			part: "red+h",
			base: normalFG,
			want: "91",
		},
		{
			name: "dimmed named color (foreground)",
			part: "red+d",
			base: normalFG,
			want: "2;31",
		},
		{
			name: "dimmed named color strikethrough (foreground)",
			part: "red+ds",
			base: normalFG,
			want: "2;9;31",
		},
		{
			name: "bold blinking named color (foreground)",
			part: "red+bB",
			base: normalFG,
			want: "1;5;31",
		},
		{
			name: "inverted named color (foreground)",
			part: "red+i",
			base: normalFG,
			want: "7;31",
		},
		{
			name: "256 color (foreground)",
			part: "160",
			base: normalFG,
			want: "38;5;160",
		},
		{
			name: "underlined 256 color (foreground)",
			part: "160+u",
			base: normalFG,
			want: "4;38;5;160",
		},
		{
			name: "truecolor (foreground)",
			part: "#ff0088",
			base: normalFG,
			want: "38;2;255;0;136",
		},
		{
			name: "named color (background)",
			part: "red",
			base: normalBG,
			want: "41",
		},
		{
			name: "light named color (background)",
			part: "red+h",
			base: normalBG,
			want: "101",
		},
		{
			name: "dimmed named color (background)",
			part: "red+d",
			base: normalBG,
			want: "2;41",
		},
		{
			name: "dimmed named color strikethrough (background)",
			part: "red+ds",
			base: normalBG,
			want: "2;9;41",
		},
		{
			name: "bold blinking named color (background)",
			part: "red+bB",
			base: normalBG,
			want: "1;5;41",
		},
		{
			name: "inverted named color (background)",
			part: "red+i",
			base: normalBG,
			want: "7;41",
		},
		{
			name: "256 color (background)",
			part: "160",
			base: normalBG,
			want: "48;5;160",
		},
		{
			name: "underlined 256 color (background)",
			part: "160+u",
			base: normalBG,
			want: "4;48;5;160",
		},
		{
			name: "truecolor (background)",
			part: "#ff0088",
			base: normalBG,
			want: "48;2;255;0;136",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			colorPartCode(buf, tt.part, tt.base)
			if got := buf.String(); got != tt.want {
				t.Fatalf("colorPartCode() = %q, expected %q", got, tt.want)
			}
		})
	}
}

func TestRgbCode(t *testing.T) {
	buf := &bytes.Buffer{}
	rgbCode(buf, "4488cc", normalFG)

	want := "38;2;68;136;204"
	if got := buf.String(); got != want {
		t.Fatalf("rgbCode() = %q, expected %q", got, want)
	}
}

func alwaysTTY() bool {
	return true
}
