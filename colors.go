package captainslog

import (
	"fmt"

	"github.com/fatih/color"
)

// Color print functions
var (
	purple = color.New(color.FgMagenta).PrintfFunc()
	blue   = color.New(color.FgBlue).PrintfFunc()
	green  = color.New(color.FgGreen).PrintfFunc()
	yellow = color.New(color.FgYellow).PrintfFunc()
	red    = color.New(color.FgRed).PrintfFunc()
)

func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
