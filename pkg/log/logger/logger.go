package logger

import (
	"io"
)

type Logger interface {
	SetLevel(string)
	GetLevel() string
	SetStream(io.Writer)
	GetStream() io.Writer
	Print(...interface{})
	Println(...interface{})
	Printf(string, ...interface{})
	Panic(...interface{})
	Panicln(...interface{})
	Panicf(string, ...interface{})
	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})
}
