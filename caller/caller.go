package caller

import (
	"runtime"
	"strings"
)

// name for anonymous calling functions
const anonymous = "anon"

// GetName returns the name of the n-th caller up the stack
func GetName(skip int) string {
	callers := make([]uintptr, skip+3)
	n := runtime.Callers(0, callers)
	target := runtime.Frame{
		Function: anonymous,
	}
	if n > 0 {
		frames := runtime.CallersFrames(callers)
		index := 0
		for current, more := frames.Next(); more && index < n+skip; current, more = frames.Next() {
			target = current
			index++
		}
	}
	return target.Function
}

// Shorten returns a shorter version of a full path name
func Shorten(path string, maxLen int) string {
	pkg := getLastPart(path, "/")
	if len(pkg) < maxLen {
		return pkg
	}
	method := getLastPart(pkg, ".")
	if len(method) < maxLen {
		return method
	}
	return method[:maxLen-2] + ".."
}

func getLastPart(text string, sep string) string {
	parts := strings.Split(text, sep)
	return parts[len(parts)-1]
}
