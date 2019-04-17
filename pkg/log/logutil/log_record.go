package logutil

// Contains LogRecord entity and method for log level setting.

// LogRecord is entity for logging.
// JSON tags allows to make logs in JSON format.
type LogRecord struct {
	// LogLevel is a field for log level.
	LogLevel string `json:"level"`

	// LogTime is a field for log time.
	LogTime string `json:"time"`

	// CallPlace is a field for call place.
	CallPlace string `json:"call_place"`

	// Msg is a field for log message.
	Msg string `json:"message"`
}

// SetLevel sets level of log for record.
func (lr *LogRecord) SetLevel(level string) {
	lr.LogLevel = level
}
