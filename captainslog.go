package captainslog

var defaultLogger *Logger

func getDefaultLogger() *Logger {
	if defaultLogger == nil {
		defaultLogger = NewLogger()
		defaultLogger.callerDepth = 6
	}
	return defaultLogger
}

// SetTimeFormat sets the time format for the default logger
func SetTimeFormat(timeFormat string) {
	getDefaultLogger().TimeFormat = timeFormat
}

// SetName overrides the caller name for the default logger
func SetName(name string) {
	getDefaultLogger().Name = name
}

// SetLevel sets the logging level for the default logger
func SetLevel(level Level) {
	getDefaultLogger().Level = level
}

// Trace logs trace-level messages with the default logger
func Trace(format string, args ...interface{}) {
	getDefaultLogger().Trace(format, args...)
}

// Debug logs debug-level messages with the default logger
func Debug(format string, args ...interface{}) {
	getDefaultLogger().Debug(format, args...)
}

// Info logs info-level messages with the default logger
func Info(format string, args ...interface{}) {
	getDefaultLogger().Info(format, args...)
}

// Warn logs warning-level messages with the default logger
func Warn(format string, args ...interface{}) {
	getDefaultLogger().Warn(format, args...)
}

// Error logs error-level messages with the default logger
func Error(format string, args ...interface{}) {
	getDefaultLogger().Error(format, args...)
}
