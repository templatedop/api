package log

import (
	"strings"

	"github.com/rs/zerolog"
)


func FetchLogLevel(level string) zerolog.Level {
	switch level {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "no-level":
		return zerolog.NoLevel
	case "disabled":
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}


type LogOutputWriter int

const (
	StdoutOutputWriter LogOutputWriter = iota
	NoopOutputWriter
	TestOutputWriter
	ConsoleOutputWriter
)


func (l LogOutputWriter) String() string {
	switch l {
	case NoopOutputWriter:
		return Noop
	case TestOutputWriter:
		return Test
	case ConsoleOutputWriter:
		return Console
	default:
		return Stdout
	}
}

// FetchLogOutputWriter returns a [LogOutputWriter] for a given value.
func FetchLogOutputWriter(l string) LogOutputWriter {
	switch strings.ToLower(l) {
	case Noop:
		return NoopOutputWriter
	case Test:
		return TestOutputWriter
	case Console:
		return ConsoleOutputWriter
	default:
		return StdoutOutputWriter
	}
}
