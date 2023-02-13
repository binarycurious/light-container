package telemetry

import "strings"

// LogLevel type for logging
type LogLevel byte

const (
	//LogLevelDebug - 0
	LogLevelDebug = iota
	//LogLevelInfo - 1
	LogLevelInfo
	//LogLevelWarn - 2
	LogLevelWarn
	//LogLevelError - 3
	LogLevelError
	//LogLevelNone - 4
	LogLevelNone
)

// GetLogLevel - returns the loglevel for the given string
func GetLogLevel(loglvl string) LogLevel {
	switch strings.ToLower(loglvl) {
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	case "none":
		return LogLevelNone
	default:
		return LogLevelInfo
	}
}

// Logger interface for generic app logging
type Logger interface {
	Log(msg string)
	LogDebug(msg string)
	LogWarn(msg string)
	LogError(msg string)
	LogFatal(msg string)
	SetActiveLogLevel(LogLevel)
}
