package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	AppEnvProd        = "prod"    // prod environment
	AppEnvDev         = "dev"     // dev environment
	AppEnvTest        = "test"    // test environment
	DefaultAppName    = "app"     // default application name
	DefaultAppVersion = "unknown" // default application version
)

type Config struct {
	*viper.Viper
	//Subconfig *SubConfig
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		v,
		//NewSubConfig(v),
	}
}

type SubConfig struct {
	v *viper.Viper
}

func NewSubConfig(v *viper.Viper) *SubConfig {
	return &SubConfig{v: v}
}

func (sc *SubConfig) GetString(key string) string {
	return sc.v.GetString(key)
}

func (sc *SubConfig) GetInt(key string) int {
	return sc.v.GetInt(key)
}

func (sc *SubConfig) GetBool(k string) bool {
	return sc.v.GetBool(k)
}

func (sc *SubConfig) GetDuration(k string) time.Duration {
	return sc.v.GetDuration(k)
}

func (sc *SubConfig) GetFloat64(k string) float64 {
	return sc.v.GetFloat64(k)
}

func (sc *SubConfig) GetInt32(k string) int32 {
	return sc.v.GetInt32(k)

}

func (sc *SubConfig) GetInt64(k string) int64 {
	return sc.v.GetInt64(k)

}

func (sc *SubConfig) GetIntSlice(k string) []int {
	return sc.v.GetIntSlice(k)

}

func (sc *SubConfig) GetStringMap(k string) map[string]any {
	return sc.v.GetStringMap(k)

}

func (sc *SubConfig) GetStringSlice(k string) []string {
	return sc.v.GetStringSlice(k)

}

func (sc *SubConfig) GetTime(k string) time.Time {
	return sc.v.GetTime(k)

}

func (sc *SubConfig) Get(key string) interface{} {
	return sc.v.Get(key)
}
func (sc *SubConfig) Set(key string, value interface{}) {
	sc.v.Set(key, value)
}

// func (sc *SubConfig) SetDefault(key string, value interface{}) {
// 	sc.v.SetDefault(key, value)
// }

func (sc *SubConfig) Exists(key string) bool {
	return sc.v.IsSet(key)
}
func (sc *SubConfig) Of(root string) *SubConfig {
	subViper := sc.v.Sub(root) // Create a new Viper instance for the sub-configuration
	if subViper == nil {
		//panic(fmt.Sprintf("no sub-config found for root key: %s", root))
	}
	return &SubConfig{v: subViper}
}

func (c *Config) Of(section string) *SubConfig {
	subViper := c.Sub(section)
	if subViper == nil {
		//panic(fmt.Sprintf("no sub-config found for root key: %s", root))
	}
	return &SubConfig{v: subViper}
	// return c.GetStringMap(section)
}

func ToStruct[T any](v *Config, root string, cfgStruct *T) error {
	subViper := v.Sub(root)
	if subViper == nil {
		return fmt.Errorf("no sub-config found for root key: %s", root)
	}
	return subViper.Unmarshal(cfgStruct)
}

func ToSubStruct[T any](v *viper.Viper, cfgStruct *T) error {
	return v.Unmarshal(cfgStruct)
}

func (c *Config) GetEnvVar(envVar string) string {
	return os.Getenv(envVar)
}

func (c *Config) AppName() string {
	return c.GetString("app.name")
}

func (c *Config) AppEnv() string {
	return c.GetString("app.env")
}

func (c *Config) AppVersion() string {
	return c.GetString("app.version")
}

func (c *Config) AppDebug() bool {
	return c.GetBool("app.debug")
}

func (c *Config) IsProdEnv() bool {
	return c.AppEnv() == AppEnvProd
}

func (c *Config) IsDevEnv() bool {
	return c.AppEnv() == AppEnvDev
}

func (c *Config) IsTestEnv() bool {
	return c.AppEnv() == AppEnvTest
}
