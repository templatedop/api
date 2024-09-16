package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ConfigFactory interface {
	Create(options ...ConfigOption) (*Config, error)
}

type DefaultConfigFactory struct{}

func NewDefaultConfigFactory() ConfigFactory {
	return &DefaultConfigFactory{}
}

func (f *DefaultConfigFactory) Create(options ...ConfigOption) (*Config, error) {
	appliedOptions := DefaultConfigOptions()
	for _, opt := range options {
		opt(&appliedOptions)
	}

	v := viper.New()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigName(appliedOptions.FileName)
	for _, path := range appliedOptions.FilePaths {
		v.AddConfigPath(path)
	}

	f.setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv != "" {
		v.SetConfigName(fmt.Sprintf("%s.%s", appliedOptions.FileName, appEnv))
		if err := v.MergeInConfig(); err != nil {
			if errors.As(err, &viper.ConfigFileNotFoundError{}) {
				return nil, fmt.Errorf("could not load config file for env %s: %w", appEnv, err)
			} else {
				return nil, fmt.Errorf("could not merge config for env %s: %w", appEnv, err)
			}
		}
	}

	for _, key := range v.AllKeys() {
		val := v.GetString(key)
		if strings.Contains(val, "${") {
			v.Set(key, os.ExpandEnv(val))
		}
	}
	
	return NewConfig(v), nil
}

func (f *DefaultConfigFactory) setDefaults(v *viper.Viper) {
	v.SetDefault("app.name", DefaultAppName)
	v.SetDefault("app.version", DefaultAppVersion)
	v.SetDefault("app.debug", false)
}
