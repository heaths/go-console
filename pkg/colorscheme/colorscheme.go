package colorscheme

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white

	csi   = "\x1b["
	sgr   = "m"
	reset = "\x1b[0m"

	normal        = "0;"
	bold          = "1;"
	dim           = "2;"
	underline     = "4;"
	blink         = "5;"
	invert        = "7;"
	strikethrough = "9;"

	color256 = "5;"
	colorRGB = "2;"

	colorOffset = 8
	normalFG    = 30
	normalBG    = 40
	lightOffset = 60
)

var (
	indexedColors = map[string]int{
		"black":   black,
		"red":     red,
		"green":   green,
		"yellow":  yellow,
		"blue":    blue,
		"magenta": magenta,
		"cyan":    cyan,
		"white":   white,
	}
)

// ColorScheme formats text with different colors and styles.
type ColorScheme struct {
	colors map[string]func(string) string
	isTTY  func() bool
}

type ColorSchemeOption func(*ColorScheme)

// New creates a new ColorScheme with options like WithTTY.
func New(opts ...ColorSchemeOption) *ColorScheme {
	return &ColorScheme{
		colors: make(map[string]func(string) string),
	}
}

// ColorFunc returns a function to format text with a given style. The resulting
// function is cached to improve performance with subsequent use.
func (cs *ColorScheme) ColorFunc(style string) func(string) string {
	if fn, ok := cs.colors[style]; ok {
		return fn
	}

	if style == "" {
		return func(s string) string {
			return s
		}
	}

	buf := colorCode(style)
	fn := func(s string) string {
		buf.WriteString(s)
		buf.WriteString(reset)
		return buf.String()
	}

	cs.colors[style] = fn
	return fn
}

// WithTTY sets a function for ColorScheme to determine if the target Writer
// represents a TTY and avoid writing terminal sequences.
func WithTTY(isTTY func() bool) ColorSchemeOption {
	return func(cs *ColorScheme) {
		cs.isTTY = isTTY
	}
}

// colorCode is compatible with github.com/mgutz/ansi with truecolor support.
func colorCode(style string) *bytes.Buffer {
	buf := &bytes.Buffer{}

	switch {
	case style == "" || style == "off":
		return buf

	case style == "reset":
		buf.WriteString(reset)
		return buf
	}

	styles := strings.Split(style, ":")
	stylesLength := len(styles)

	// Write CSI and reset.
	buf.WriteString(csi)
	buf.WriteString(normal)

	// Write foreground.
	colorPartCode(buf, styles[0], normalFG)

	// Write background.
	if stylesLength > 1 {
		// Only write separator if we wrote the foreground.
		if len(styles[0]) > 0 {
			buf.WriteRune(';')
		}
		colorPartCode(buf, styles[1], normalBG)
	}
	buf.WriteString(sgr)

	return buf
}

func colorPartCode(buf *bytes.Buffer, part string, base int) {
	if part == "" {
		return
	}

	styles := strings.Split(part, "+")
	if styles[0] == "" {
		return
	}

	color, style := styles[0], ""
	if len(styles) > 1 {
		style = styles[1]
	}

	if strings.Contains(style, "b") {
		buf.WriteString(bold)
	}
	if strings.Contains(style, "d") {
		buf.WriteString(dim)
	}
	if strings.Contains(style, "B") {
		buf.WriteString(blink)
	}
	if strings.Contains(style, "u") {
		buf.WriteString(underline)
	}
	if strings.Contains(style, "i") {
		buf.WriteString(invert)
	}
	if strings.Contains(style, "s") {
		buf.WriteString(strikethrough)
	}
	if strings.Contains(style, "h") {
		base += lightOffset
	}

	if strings.HasPrefix(color, "#") && len(color) == 7 {
		rgbCode(buf, color[1:], base)
	} else if i, ok := indexedColors[color]; ok {
		buf.WriteString(strconv.Itoa(base + i))
	} else if i, err := strconv.Atoi(color); err == nil && i >= 1 && i < 256 {
		fmt.Fprintf(buf, "%d;", base+colorOffset)
		buf.WriteString(color256)
		buf.WriteString(color)
	} else {
		// Reset.
		buf.WriteRune('0')
	}
}

func rgbCode(buf *bytes.Buffer, rgb string, base int) {
	r, _ := strconv.ParseInt(rgb[0:2], 16, 64)
	g, _ := strconv.ParseInt(rgb[2:4], 16, 64)
	b, _ := strconv.ParseInt(rgb[4:6], 16, 64)

	fmt.Fprintf(buf, "%d;", base+colorOffset)
	buf.WriteString(colorRGB)
	fmt.Fprintf(buf, "%d;%d;%d", r, g, b)
}

func Black(s string) string {
	return foreground(normalFG+black, s)
}

func Red(s string) string {
	return foreground(normalFG+red, s)
}

func Green(s string) string {
	return foreground(normalFG+green, s)
}

func Yellow(s string) string {
	return foreground(normalFG+yellow, s)
}

func Blue(s string) string {
	return foreground(normalFG+blue, s)
}

func Magenta(s string) string {
	return foreground(normalFG+magenta, s)
}

func Cyan(s string) string {
	return foreground(normalFG+cyan, s)
}

func White(s string) string {
	return foreground(normalFG+white, s)
}
func LightBlack(s string) string {
	return foreground(normalFG+lightOffset+black, s)
}

func LightRed(s string) string {
	return foreground(normalFG+lightOffset+red, s)
}

func LightGreen(s string) string {
	return foreground(normalFG+lightOffset+green, s)
}

func LightYellow(s string) string {
	return foreground(normalFG+lightOffset+yellow, s)
}

func LightBlue(s string) string {
	return foreground(normalFG+lightOffset+blue, s)
}

func LightMagenta(s string) string {
	return foreground(normalFG+lightOffset+magenta, s)
}

func LightCyan(s string) string {
	return foreground(normalFG+lightOffset+cyan, s)
}

func LightWhite(s string) string {
	return foreground(normalFG+lightOffset+white, s)
}

func foreground(c int, s string) string {
	return csi + normal + strconv.Itoa(c) + "m" + s + reset
}
