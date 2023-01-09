package container

import (
	"fmt"
	"log"

	"github.com/binarycurious/light-container/telemetry"
)

// GlobalLogger impl
type GlobalLogger struct {
	hardfail bool
	loglevel telemetry.LogLevel
}

// Log (log *string) //info level logging
func (l *GlobalLogger) Log(msg *string) {
	if l.loglevel <= telemetry.LogLevelInfo {
		fmt.Printf("INFO: %s", *msg)
	}
}

// LogDebug (msg *string)
func (l *GlobalLogger) LogDebug(msg *string) {
	if l.loglevel <= telemetry.LogLevelDebug {
		fmt.Printf("DEBUG: %s", *msg)
	}
}

// LogWarn (msg *string)
func (l *GlobalLogger) LogWarn(msg *string) {
	if l.loglevel <= telemetry.LogLevelWarn {
		fmt.Printf("WARN: %s", *msg)
	}
}

// LogError (msg *string)
func (l *GlobalLogger) LogError(msg *string) {
	if l.loglevel <= telemetry.LogLevelError {
		fmt.Printf("ERROR: %s", *msg)
	}
}

// LogFatal (msg *string, hardfail)
func (l *GlobalLogger) LogFatal(msg *string) {
	if !l.hardfail {
		log.Print(*msg)
		return
	}
	log.Fatal(*msg)
}

// SetActiveLogLevel (LogLevel)
func (l *GlobalLogger) SetActiveLogLevel(lvl telemetry.LogLevel) {
	l.loglevel = lvl
}
