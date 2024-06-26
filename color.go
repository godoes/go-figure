package figure

import (
	"github.com/fatih/color"
)

var (
	ColorReset  = "reset"
	ColorRed    = "red"
	ColorGreen  = "green"
	ColorYellow = "yellow"
	ColorBlue   = "blue"
	ColorPurple = "purple"
	ColorCyan   = "cyan"
	ColorGray   = "gray"
	ColorWhite  = "white"

	colorColors map[string]color.Attribute
)

var colors = map[string]string{
	ColorReset:  "\033[0m",
	ColorRed:    "\033[31m",
	ColorGreen:  "\033[32m",
	ColorYellow: "\033[33m",
	ColorBlue:   "\033[34m",
	ColorPurple: "\033[35m",
	ColorCyan:   "\033[36m",
	ColorGray:   "\033[37m",
	ColorWhite:  "\033[97m",
}
