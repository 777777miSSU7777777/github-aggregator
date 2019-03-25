package loggerfactory

import (
	"io"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logger"
)

func NewJson(level string, stream io.Writer) *logger.JsonLogger {
	jsonLogger := &logger.JsonLogger{}
	jsonLogger.SetLevel(level)
	jsonLogger.SetStream(stream)
	return jsonLogger
}
