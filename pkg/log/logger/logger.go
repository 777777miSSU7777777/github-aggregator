// Package logger contains an interface for loggers.
package logger

import (
	"io"
)

// Logger is an interface for loggers.
type Logger interface {
	// SetLevel sets level of logging.
	SetLevel(string)

	// GetLevel returns level of logging.
	GetLevel() string

	// SetStream sets logging output stream.
	SetStream(io.Writer)

	// GetStream returns logging output stream.
	GetStream() io.Writer

	// Print prints log record.
	Print(...interface{})

	// Println prints log record with new line.
	Println(...interface{})

	// Printf print log record with formatted message.
	Printf(string, ...interface{})

	// Panic prints log record and then panics.
	Panic(...interface{})

	// Panicln prints log record with new line and then panics.
	Panicln(...interface{})

	// Panicf prints log record with formatted message and the panics.
	Panicf(string, ...interface{})

	// Fatal prints log record and then makes os.Exit(1).
	Fatal(...interface{})

	// Fatalln prints log record with new line and then makes os.Exit(1).
	Fatalln(...interface{})

	// Fatalf prints log record with formatted message and then makes os.Exit(1).
	Fatalf(string, ...interface{})
}
