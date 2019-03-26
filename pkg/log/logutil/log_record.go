package logutil

// LogRecord is entity for logging.
// JSON tags allows to make logs in JSON format.
type LogRecord struct {
	// LogLevel is a field for log level.
	LogLevel string `json:"level"`

	// LogTime is a field for log time.
	LogTime string `json:"time"`

	// FileShortcut is a field for file shortcut.
	FileShortcut string `json:"file_shortcut"`

	// Msg is a field for log message.
	Msg string `json:"message"`
}

// SetLevel sets level of log for record.
func (lr *LogRecord) SetLevel(level string) {
	lr.LogLevel = level
}