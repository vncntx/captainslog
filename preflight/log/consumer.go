package log

import "os"

// LogsConsumer is a function that consumes two streams
type LogsConsumer func(stdout *os.File, stderr *os.File)
