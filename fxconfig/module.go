package fxconfig

import (
	"os"

	"github.com/templatedop/api/config"
	"github.com/templatedop/api/module"
	"go.uber.org/fx"
)
const ModuleName = "config"
func ConfigModule() (*module.Module, error) {
	m := module.New(ModuleName)
	m.Provide(config.NewDefaultConfigFactory,
		NewFxConfig,)

	return m, nil
}


type FxConfigParam struct {
	fx.In
	Factory config.ConfigFactory
}

func NewFxConfig(p FxConfigParam) (*config.Config, error) {
	return p.Factory.Create(
		config.WithFileName("config"),
		config.WithFilePaths(
			".",
			"./configs",
			os.Getenv("APP_CONFIG_PATH"),
		),
	)
}
