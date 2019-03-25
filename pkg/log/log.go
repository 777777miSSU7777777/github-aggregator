package log

import (
	"io/ioutil"
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logger"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/loggerfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logutil"
)

var Info logger.Logger

var Warning logger.Logger

var Trace logger.Logger

var Error logger.Logger

func init() {
	logfactory.SetLogDepth(logutil.DefaultLogDepth)

	Info = loggerfactory.NewJson(logutil.INFO, os.Stdout)

	Warning = loggerfactory.NewJson(logutil.WARNING, os.Stdout)

	Trace = loggerfactory.NewJson(logutil.TRACE, ioutil.Discard)

	Error = loggerfactory.NewJson(logutil.ERROR, os.Stderr)
}
