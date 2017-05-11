package log

import (
	"github.com/go-log/log"
	golog "github.com/go-log/log/log"
)

var (
	// the local logger
	logger log.Logger = golog.New()
)

// Log makes use of github.com/go-log/log.Log
func Log(v ...interface{}) {
	logger.Log(v...)
}

// Logf makes use of github.com/go-log/log.Logf
func Logf(format string, v ...interface{}) {
	logger.Logf(format, v...)
}

// SetLogger sets the local logger
func SetLogger(l log.Logger) {
	logger = l
}
