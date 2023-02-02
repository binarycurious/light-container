package telemetry

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

// Logger interface for generic app logging
type Logger interface {
	Log(msg string)
	LogDebug(msg string)
	LogWarn(msg string)
	LogError(msg string)
	LogFatal(msg string)
	SetActiveLogLevel(LogLevel)
}
