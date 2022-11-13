package ansi

const (
	ESC = "\x1b"
	CSI = ESC + "["
	ST  = ESC + "\\"
	BEL = "\a"

	// Select graphics rendition (SGR) codes.
	Reset = CSI + "0m"
)
