package loggerfactory

// A factory for JSONLoggers.

import (
	"io"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logger"
)

// NewJSON returns an instance of JSONLogger.
func NewJSON(level string, stream io.Writer) *logger.JSONLogger {
	JSONLogger := &logger.JSONLogger{}
	JSONLogger.SetLevel(level)
	JSONLogger.SetStream(stream)
	return JSONLogger
}
