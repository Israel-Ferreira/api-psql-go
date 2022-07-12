package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type ApiConfig struct {
	Port string
}

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type config struct {
	API ApiConfig
	DB  DbConfig
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = ApiConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DbConfig{

		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.name"),
	}

	return nil
}

func GetDb() DbConfig {
	return cfg.DB
}

func GetApiPort() string {
	return cfg.API.Port
}
