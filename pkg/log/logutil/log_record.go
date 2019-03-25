package logutil

type LogRecord struct {
	LogLevel     string `json:"level"`
	LogTime      string `json:"time"`
	FileShortcut string `json:"file_shortcut"`
	Msg          string `json:"message"`
}

func (lr *LogRecord) SetLevel(level string) {
	lr.LogLevel = level
}
