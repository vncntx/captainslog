package msg

import (
	"github.com/fatih/color"
)

// Color adds color codes to a string
type Color func(string, ...interface{}) string

// Color print functions
var (
	cyan   = color.New(color.FgHiCyan).SprintfFunc()
	blue   = color.New(color.FgHiBlue).SprintfFunc()
	green  = color.New(color.FgHiGreen).SprintfFunc()
	yellow = color.New(color.FgHiYellow).SprintfFunc()
	red    = color.New(color.FgHiRed).SprintfFunc()
)
