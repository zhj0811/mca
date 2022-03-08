package zlogging

import (
	"io"

	"go.uber.org/zap"
)

var Global *Logging

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
}

// Init initializes logging with the provided config.
func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}

func SetWriter(w io.Writer) io.Writer {
	return Global.SetWriter(w)
}

func LoggerLevel() string {
	return Global.level.String()
}

// MustGetLogger creates a logger with the specified name. If an invalid name
// is provided, the operation will panic.
func MustGetLogger(loggerName string) *zap.SugaredLogger {
	return Global.Logger(loggerName)
}

func SetLogLevel(level string) {
	Global.level = NameToLevel(level)
	return
}
