package log

import (
	"github.com/rs/zerolog"
	"sync"
)

var once sync.Once


type LoggerFactory interface {
	Create(options ...LoggerOption) (*Logger, error)
}

type DefaultLoggerFactory struct{}

func NewDefaultLoggerFactory() LoggerFactory {
	return &DefaultLoggerFactory{}
}


func (f *DefaultLoggerFactory) Create(options ...LoggerOption) (*Logger, error) {
	appliedOpts := DefaultLoggerOptions()
	for _, applyOpt := range options {
		applyOpt(&appliedOpts)
	}

	logger := zerolog.
		New(appliedOpts.OutputWriter).
		With().		
		Timestamp().
		Str(Service, appliedOpts.ServiceName).
		Logger().
		Level(appliedOpts.Level)

	once.Do(func() {
		zerolog.DefaultContextLogger = &logger
	})

	return &Logger{&logger}, nil
}
