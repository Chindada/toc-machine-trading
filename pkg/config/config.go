// Package config package config
package config

import (
	"sync"

	"toc-machine-trading/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	globalConfig *Config
	once         sync.Once
)

// GetConfig -.
func GetConfig() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}

	once.Do(parseConfigFile)
	return globalConfig, nil
}

func parseConfigFile() {
	newConfig := Config{}
	err := cleanenv.ReadConfig("./configs/config.yml", &newConfig)
	if err != nil {
		logger.Get().Panic(err)
	}

	err = cleanenv.ReadEnv(&newConfig)
	if err != nil {
		logger.Get().Panic(err)
	}

	globalConfig = &newConfig
}

// Config -.
type Config struct {
	HTTP     `env-required:"true" yaml:"http"`
	Postgres `env-required:"true" yaml:"postgres"`
	Sinopac  `env-required:"true" yaml:"sinopac"`

	Deployment string `env-required:"true" env:"DEPLOYMENT"`
}

// HTTP -.
type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

// Postgres -.
type Postgres struct {
	PoolMax int `env-required:"true" yaml:"pool_max"`

	URL    string `env-required:"true" env:"PG_URL"`
	DBName string `env-required:"true" env:"DB_NAME"`
}

// Sinopac -.
type Sinopac struct {
	PoolMax int `env-required:"true" yaml:"pool_max"`

	URL string `env-required:"true" env:"SINOPAC_URL"`
}
