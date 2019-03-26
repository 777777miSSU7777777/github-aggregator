package logger

// JSONLogger is an implentation of Logger interface.
// Makes logs in JSON format.
// Provides such info as level, time, shortcut to call place and message.

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logfactory"
)

// JSONLogger is an implemtation of Logger interface.
// Makes logs in json format.
type JSONLogger struct {
	// level is a private field which responsible for level of logging.
	level string

	// stream is a private field which responsible for output stream for logging.
	stream io.Writer
}

// SetLevel sets level of logging.
func (jl *JSONLogger) SetLevel(level string) {
	jl.level = level
}

// GetLevel returns level of logging.
func (jl JSONLogger) GetLevel() string {
	return jl.level
}

// SetStream sets logging output stream.
func (jl *JSONLogger) SetStream(stream io.Writer) {
	jl.stream = stream
}

// GetStream returns logging output stream.
func (jl JSONLogger) GetStream() io.Writer {
	return jl.stream
}

// Print prints log record in json format.
func (jl JSONLogger) Print(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprint(jl.stream, string(jsonLog))
}

// Println prints log record in json format with new line.
func (jl JSONLogger) Println(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
}

// Printf print log record in json format with formatted message.
func (jl JSONLogger) Printf(format string, a ...interface{}) {
	logRecord := logfactory.NewF(format, a...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
}

// Panic prints log record in json format and then panics.
func (jl JSONLogger) Panic(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprint(jl.stream, string(jsonLog))
	panic(logRecord.Msg)
}

// Panicln prints log record in json format with new line and then panics.
func (jl JSONLogger) Panicln(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	panic(logRecord.Msg)
}

// Panicf prints log record in json format with formatted message and the panics.
func (jl JSONLogger) Panicf(format string, a ...interface{}) {
	logRecord := logfactory.NewF(format, a...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	panic(logRecord.Msg)
}

// Fatal prints log record in json format and then makes os.Exit(1).
func (jl JSONLogger) Fatal(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprint(jl.stream, string(jsonLog))
	os.Exit(1)
}

// Fatalln prints log record in json format with new line and then makes os.Exit(1).
func (jl JSONLogger) Fatalln(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	os.Exit(1)
}

// Fatalf prints log record in json format with formatted message and then makes os.Exit(1).
func (jl JSONLogger) Fatalf(format string, a ...interface{}) {
	logRecord := logfactory.NewF(format, a...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	os.Exit(1)
}
