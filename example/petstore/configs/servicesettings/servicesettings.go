package servicesettings

import (
	"os"
	"sync"
)

type Config struct {
	SysName	string
}

var config *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		
		config = &Config{
			SysName: os.Getenv("SERVICE_NAME"),
		}
	})
	return config
}
