package colorscheme

import (
	"strconv"
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
	defaultColor = 9

	csi   = "\x1b["
	reset = "\x1b[0m"

	normal        = "0;"
	bold          = "1;"
	dim           = "2;"
	italic        = "3;"
	underline     = "4;"
	blink         = "5;"
	invert        = "7;"
	strikethrough = "9;"

	normalFG = 30
	normalBG = 40
	lightFG  = 90
	lightBG  = 100
)

type ColorScheme struct {
}

func NewColorScheme() *ColorScheme {
	return &ColorScheme{}
}

func (cs *ColorScheme) ColorFunc(color string) func(string) string {
	return func(s string) string {
		return s
	}
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
	return foreground(lightFG+black, s)
}

func LightRed(s string) string {
	return foreground(lightFG+red, s)
}

func LightGreen(s string) string {
	return foreground(lightFG+green, s)
}

func LightYellow(s string) string {
	return foreground(lightFG+yellow, s)
}

func LightBlue(s string) string {
	return foreground(lightFG+blue, s)
}

func LightMagenta(s string) string {
	return foreground(lightFG+magenta, s)
}

func LightCyan(s string) string {
	return foreground(lightFG+cyan, s)
}

func LightWhite(s string) string {
	return foreground(lightFG+white, s)
}

func foreground(c int, s string) string {
	return csi + normal + strconv.Itoa(c) + "m" + s + reset
}
