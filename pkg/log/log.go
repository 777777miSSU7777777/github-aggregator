package log

import (
	"io/ioutil"
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logger"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/loggerfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logutil"
)

// Info logger for info level of logging.
var Info logger.Logger

// Warning logger for warning level of logging.
var Warning logger.Logger

// Trace logger for trace level of logging.
var Trace logger.Logger

// Error logger for error level of logging.
var Error logger.Logger

func init() {
	logfactory.SetLogDepth(logutil.DefaultLogDepth)

	Info = loggerfactory.NewJSON(logutil.INFO, os.Stdout)

	Warning = loggerfactory.NewJSON(logutil.WARNING, os.Stdout)

	Trace = loggerfactory.NewJSON(logutil.TRACE, ioutil.Discard)

	Error = loggerfactory.NewJSON(logutil.ERROR, os.Stderr)
}
