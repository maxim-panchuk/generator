package config

import (
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	PathToSwaggerFile    string
	PathToRepositoryRoot string
	ServiceName          string
}

var config *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		viper.AutomaticEnv()
		config = &Config{
			PathToSwaggerFile:    viper.GetString("SWAGGER_FILE"),
			PathToRepositoryRoot: viper.GetString("ROOT_PATH"),
			ServiceName:          viper.GetString("SERVICE_NAME"),
		}
	})
	return config
}
