package caller

import (
	"runtime"
)

// name for anonymous calling functions
const anonymous = "anonymous"

// GetName returns the name of the n-th caller up the stack
func GetName(n int) string {
	callerPtrs := make([]uintptr, 1)
	callersCount := runtime.Callers(n, callerPtrs)
	if callersCount == 0 {
		return anonymous
	}
	callerFunc := runtime.FuncForPC(callerPtrs[0])
	if callerFunc == nil {
		return anonymous
	}
	return callerFunc.Name()
}
