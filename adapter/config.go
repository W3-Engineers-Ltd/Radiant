package adapter

import (
	"github.com/W3-Engineers-Ltd/Radiant/adapter/session"
	newCfg "github.com/W3-Engineers-Ltd/Radiant/core/config"
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

// Config is the main struct for BConfig
type Config web.Config

// Listen holds for http and https related config
type Listen web.Listen

// WebConfig holds web related config
type WebConfig web.WebConfig

// SessionConfig holds session related config
type SessionConfig web.SessionConfig

// LogConfig holds Log related config
type LogConfig web.LogConfig

var (
	// BConfig is the default config for Application
	BConfig *Config
	// AppConfig is the instance of Config, store the config information from file
	AppConfig *radiantAppConfig
	// AppPath is the absolute path to the app
	AppPath string
	// GlobalSessions is the instance for the session manager
	GlobalSessions *session.Manager

	// appConfigPath is the path to the config files
	appConfigPath string
	// appConfigProvider is the provider for the config, default is ini
	appConfigProvider = "ini"
	// WorkPath is the absolute path to project root directory
	WorkPath string
)

func init() {
	BConfig = (*Config)(web.BConfig)
	AppPath = web.AppPath

	WorkPath = web.WorkPath

	AppConfig = &radiantAppConfig{innerConfig: (newCfg.Configer)(web.AppConfig)}
}

// LoadAppConfig allow developer to apply a config file
func LoadAppConfig(adapterName, configPath string) error {
	return web.LoadAppConfig(adapterName, configPath)
}

type radiantAppConfig struct {
	innerConfig newCfg.Configer
}

func (b *radiantAppConfig) Set(key, val string) error {
	if err := b.innerConfig.Set(BConfig.RunMode+"::"+key, val); err != nil {
		return b.innerConfig.Set(key, val)
	}
	return nil
}

func (b *radiantAppConfig) String(key string) string {
	if v, err := b.innerConfig.String(BConfig.RunMode + "::" + key); v != "" && err != nil {
		return v
	}
	res, _ := b.innerConfig.String(key)
	return res
}

func (b *radiantAppConfig) Strings(key string) []string {
	if v, err := b.innerConfig.Strings(BConfig.RunMode + "::" + key); len(v) > 0 && err != nil {
		return v
	}
	res, _ := b.innerConfig.Strings(key)
	return res
}

func (b *radiantAppConfig) Int(key string) (int, error) {
	if v, err := b.innerConfig.Int(BConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int(key)
}

func (b *radiantAppConfig) Int64(key string) (int64, error) {
	if v, err := b.innerConfig.Int64(BConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int64(key)
}

func (b *radiantAppConfig) Bool(key string) (bool, error) {
	if v, err := b.innerConfig.Bool(BConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Bool(key)
}

func (b *radiantAppConfig) Float(key string) (float64, error) {
	if v, err := b.innerConfig.Float(BConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return b.innerConfig.Float(key)
}

func (b *radiantAppConfig) DefaultString(key string, defaultVal string) string {
	if v := b.String(key); v != "" {
		return v
	}
	return defaultVal
}

func (b *radiantAppConfig) DefaultStrings(key string, defaultVal []string) []string {
	if v := b.Strings(key); len(v) != 0 {
		return v
	}
	return defaultVal
}

func (b *radiantAppConfig) DefaultInt(key string, defaultVal int) int {
	if v, err := b.Int(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *radiantAppConfig) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := b.Int64(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *radiantAppConfig) DefaultBool(key string, defaultVal bool) bool {
	if v, err := b.Bool(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *radiantAppConfig) DefaultFloat(key string, defaultVal float64) float64 {
	if v, err := b.Float(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *radiantAppConfig) DIY(key string) (interface{}, error) {
	return b.innerConfig.DIY(key)
}

func (b *radiantAppConfig) GetSection(section string) (map[string]string, error) {
	return b.innerConfig.GetSection(section)
}

func (b *radiantAppConfig) SaveConfigFile(filename string) error {
	return b.innerConfig.SaveConfigFile(filename)
}
