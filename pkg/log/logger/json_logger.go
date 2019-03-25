package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logfactory"
)

type JsonLogger struct {
	level  string
	stream io.Writer
}

func (jl *JsonLogger) SetLevel(level string) {
	jl.level = level
}

func (jl JsonLogger) GetLevel() string {
	return jl.level
}

func (jl *JsonLogger) SetStream(stream io.Writer) {
	jl.stream = stream
}

func (jl JsonLogger) GetStream() io.Writer {
	return jl.stream
}

func (jl JsonLogger) Print(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprint(jl.stream, string(jsonLog))
}

func (jl JsonLogger) Println(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
}

func (jl JsonLogger) Printf(format string, a ...interface{}) {
	logRecord := logfactory.NewF(format, a...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
}

func (jl JsonLogger) Panic(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprint(jl.stream, string(jsonLog))
	panic(logRecord.Msg)
}

func (jl JsonLogger) Panicln(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	panic(logRecord.Msg)
}

func (jl JsonLogger) Panicf(format string, a ...interface{}) {
	logRecord := logfactory.NewF(format, a...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	panic(logRecord.Msg)
}

func (jl JsonLogger) Fatal(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprint(jl.stream, string(jsonLog))
	os.Exit(1)
}

func (jl JsonLogger) Fatalln(v ...interface{}) {
	logRecord := logfactory.New(v...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	os.Exit(1)
}

func (jl JsonLogger) Fatalf(format string, a ...interface{}) {
	logRecord := logfactory.NewF(format, a...)
	logRecord.SetLevel(jl.level)
	jsonLog, _ := json.Marshal(logRecord)
	fmt.Fprintln(jl.stream, string(jsonLog))
	os.Exit(1)
}
