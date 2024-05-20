
package postgres

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Login    string
	Password string
	Database string
	Schema   string
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		setDefaults()
		viper.AutomaticEnv()
		config = &Config{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Login:    viper.GetString("DB_LOGIN"),
			Password: viper.GetString("DB_PASSWORD"),
			Database: viper.GetString("DB_DATABASE"),
			Schema:   viper.GetString("DB_SCHEMA"),
		}
	})
	return config
}

var gormOnce sync.Once
var instance *gorm.DB

func GetGorm() *gorm.DB {
	gormOnce.Do(func() {
		config := GetConfig()
		db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			config.Host, config.Login, config.Password, config.Database, config.Port)))
		if err != nil {
			panic("connection database failed")
		}
		instance = db
	})
	return instance
}

func setDefaults() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_LOGIN", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_DATABASE", "postgres")
	viper.SetDefault("DB_SCHEMA", "public")
}
