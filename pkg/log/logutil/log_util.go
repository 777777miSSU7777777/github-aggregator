package logutil

import (
	"fmt"
	"runtime"
	"time"
)

// DefaultLogDepth default depth for runtime.Caller which goes through log factory and print functions.
const DefaultLogDepth int = 4

// timeFormat is
const timeFormat string = "02-01-2006 15:04:05"

// GetFnAndLineNumber returns function and code line number where it was called.
// Int param "depth" defines depth for runtime.Caller.
func GetFnAndLineNumber(depth int) (string, int) {
	_, fn, line, _ := runtime.Caller(depth)
	return fn, line
}

// GetCurrentTime returns string reprentation of current time.
func GetCurrentTime() string {
	now := time.Now()
	return now.Format(timeFormat)
}

// GetCallPlace returns shortcut to file and code line number using GetFnAndLineNumber.
// Int param "depth" defines depth for GetFnAndLineNumber.
func GetCallPlace(depth int) string {
	fn, line := GetFnAndLineNumber(depth)
	return fmt.Sprintf("%s:%d", fn, line)
}
