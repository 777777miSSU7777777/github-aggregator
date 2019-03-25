package logutil

import (
	"fmt"
	"runtime"
	"time"
)

const DefaultLogDepth int = 4

const timeFormat string = "02-01-2006 15:04:05"

func GetFnAndLineNumber(depth int) (string, int) {
	_, fn, line, _ := runtime.Caller(depth)
	return fn, line
}

func GetCurrentTime() string {
	now := time.Now()
	return now.Format(timeFormat)
}

func GetCallPlace(depth int) string {
	fn, line := GetFnAndLineNumber(depth)
	return fmt.Sprintf("%s:%d", fn, line)
}
