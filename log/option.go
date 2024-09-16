package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Options struct {
	ServiceName  string
	Level        zerolog.Level
	OutputWriter io.Writer
}

func DefaultLoggerOptions() Options {
	return Options{
		ServiceName:  "default",
		Level:        zerolog.InfoLevel,
		OutputWriter: os.Stdout,
	}
}

type LoggerOption func(o *Options)

func WithServiceName(n string) LoggerOption {
	return func(o *Options) {
		o.ServiceName = n
	}
}

func WithLevel(l zerolog.Level) LoggerOption {
	return func(o *Options) {
		o.Level = l
	}
}

func WithOutputWriter(w io.Writer) LoggerOption {
	return func(o *Options) {
		o.OutputWriter = w
	}
}
