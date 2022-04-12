package logsymbols

import (
	"fmt"
	"os"
	"runtime"
)

type color int

const (
	black         = iota
	red     color = iota
	green   color = iota
	yellow  color = iota
	blue    color = iota
	magenta color = iota
	cyan    color = iota
	white   color = iota
	grey    color = iota
	reset   color = iota
)

var (
	// INFO is the string info symbol
	INFO string
	// SUCCESS is the string success symbol
	SUCCESS string
	// WARNING is the string warning symbol
	WARNING string
	// ERROR is the string error symbol
	ERROR string
)

func style(c color, s string) string {
	return fmt.Sprintf("\u001b[3%dm%s\u001b[3%dm", c, s, reset)
}

func init() {
	termVal, exists := os.LookupEnv("TERM")

	if (exists && termVal == "xterm-256color") || runtime.GOOS == "windows" {
		INFO = style(blue, "ℹ")
		SUCCESS = style(green, "✔")
		WARNING = style(yellow, "⚠")
		ERROR = style(red, "✖")
	} else {
		INFO = style(blue, "i")
		SUCCESS = style(green, "√")
		WARNING = style(yellow, "‼")
		ERROR = style(red, "×")
	}
}
