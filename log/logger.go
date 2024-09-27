package log

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/templatedop/api/util/appctx"
)

const (
	Level   = "level"
	Message = "message"
	Service = "service"
	Time    = "time"
	Stdout  = "stdout"
	Noop    = "noop"
	Test    = "test"
	Console = "console"
	Frame   = 3
)

var loggerCtxValuer = appctx.NewValuer[*Logger]("zerolog_logger_key").OnDefault(func() *Logger {
	existingLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{&existingLogger}
})

type Logger struct {
	logger *zerolog.Logger
}

func WithLogger(parent context.Context, l *Logger) context.Context {
	return loggerCtxValuer.Set(parent, l)
}

func (l *Logger) CallerIncluded() *Logger {
	lo := l.logger.With().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + Frame).Logger()

	return &Logger{&lo}
}

func (l *Logger) ToZerolog() *zerolog.Logger {
	return l.logger
}

func FromZerolog(logger zerolog.Logger) *Logger {
	return &Logger{&logger}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {

	l.msg(zerolog.DebugLevel, message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {

	l.log1(zerolog.InfoLevel, message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log1(zerolog.WarnLevel, message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.msg(zerolog.ErrorLevel, message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg(zerolog.FatalLevel, message, args...)

	os.Exit(1)
}

func (l *Logger) msg(level zerolog.Level, message interface{}, args ...interface{}) {

	switch msg := message.(type) {
	case error:

		l.log1(level, msg.Error(), args...)
	case string:
		l.log1(level, msg, args...)
	default:
		l.log1(zerolog.InfoLevel, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

func (l *Logger) log1(level zerolog.Level, message string, args ...interface{}) {
	lw := l.CallerIncluded()
	loggers := lw.logger.WithLevel(level)

	if len(args) == 0 {
		loggers.Msg(message)
	} else {
		loggers.Msgf(message, args...)
	}
}
