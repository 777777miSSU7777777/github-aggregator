package logfactory

import (
	"fmt"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logutil"
)

var logDepth int

// SetLogDepth sets depth for runtime.Caller depth.
func SetLogDepth(depth int) {
	logDepth = depth
}

// GetLogDepth returns depth for runtime.Caller.
func GetLogDepth() int {
	return logDepth
}

// New returns an instance of LogRecord.
func New(v ...interface{}) *logutil.LogRecord {
	s := fmt.Sprint(v...)
	logRecord := &logutil.LogRecord{LogTime: logutil.GetCurrentTime(), FileShortcut: logutil.GetCallPlace(logDepth), Msg: s}
	return logRecord
}

// NewF returns an instance of LogRecord with formatted message.
func NewF(format string, a ...interface{}) *logutil.LogRecord {
	s := fmt.Sprintf(format, a...)
	logRecord := &logutil.LogRecord{LogTime: logutil.GetCurrentTime(), FileShortcut: logutil.GetCallPlace(logDepth), Msg: s}
	return logRecord
}
