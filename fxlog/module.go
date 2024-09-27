package fxlog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/templatedop/api/config"
	"github.com/templatedop/api/log"
	"github.com/templatedop/api/module"
	"go.uber.org/fx"
)

const ModuleName = "log"

func LogModule() (*module.Module, error) {
	m := module.New(ModuleName)
	m.Provide(log.NewDefaultLoggerFactory,
		NewFxLogger)

	return m, nil
}

type FxLogParam struct {
	fx.In
	Factory log.LoggerFactory
	Config  *config.Config
}

func NewFxLogger(p FxLogParam) (*log.Logger, error) {

	var level zerolog.Level
	if p.Config.AppDebug() {
		level = zerolog.DebugLevel
	} else {
		level = log.FetchLogLevel(p.Config.GetString("log.level"))
	}
	var outputWriter io.Writer

	switch log.FetchLogOutputWriter(p.Config.GetString("log.output")) {
	case log.NoopOutputWriter:
		outputWriter = io.Discard
	case log.ConsoleOutputWriter:
		outputWriter = zerolog.ConsoleWriter{Out: os.Stderr}
	default:
		outputWriter = os.Stdout
	}

	return p.Factory.Create(
		log.WithServiceName(p.Config.AppName()),
		log.WithLevel(level),
		log.WithOutputWriter(outputWriter),
	)
}
