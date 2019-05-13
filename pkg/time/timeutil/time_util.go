package timeutil

import (
	"time"
)

const timeFormat string = "02-01-2006 15:04:05"

// GetCurrentTime returns string representation of current time.
func GetCurrentTime() string {
	now := time.Now()
	return now.Format(timeFormat)
}
