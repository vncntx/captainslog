package msg

import (
	"github.com/fatih/color"
)

// Color adds color codes to a string
type Color func(string, ...interface{}) string

// Color print functions
var (
	cyan   = color.New(color.FgCyan).SprintfFunc()
	blue   = color.New(color.FgBlue).SprintfFunc()
	green  = color.New(color.FgGreen).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	red    = color.New(color.FgRed).SprintfFunc()
)
