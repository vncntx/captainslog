package captainslog

import (
	"github.com/fatih/color"
)

// Color print functions
var (
	purple = color.New(color.FgMagenta).SprintfFunc()
	blue   = color.New(color.FgBlue).SprintfFunc()
	green  = color.New(color.FgGreen).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	red    = color.New(color.FgRed).SprintfFunc()
)
