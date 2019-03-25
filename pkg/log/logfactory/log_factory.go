package logfactory

import (
	"fmt"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logutil"
)

var logDepth int

func SetLogDepth(depth int) {
	logDepth = depth
}

func GetLogDepth() int {
	return logDepth
}

func New(v ...interface{}) *logutil.LogRecord {
	s := fmt.Sprint(v...)
	logRecord := &logutil.LogRecord{LogTime: logutil.GetCurrentTime(), FileShortcut: logutil.GetCallPlace(logDepth), Msg: s}
	return logRecord
}

func NewF(format string, a ...interface{}) *logutil.LogRecord {
	s := fmt.Sprintf(format, a...)
	logRecord := &logutil.LogRecord{LogTime: logutil.GetCurrentTime(), FileShortcut: logutil.GetCallPlace(logDepth), Msg: s}
	return logRecord
}
