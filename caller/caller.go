package caller

import (
	"runtime"
)

// name for anonymous calling functions
const anonymous = "anonymous"

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
