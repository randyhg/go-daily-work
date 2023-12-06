package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port              string      `yaml:"Port"`
	Username          string      `yaml:"Username"`
	Password          string      `yaml:"Password"`
	TokenKey          string      `yaml:"Token-Key"`
	ShowSql           bool        `yaml:"ShowSql"`
	MySqlUrl          string      `yaml:"MySqlUrl"`
	MySqlMaxIdle      int         `yaml:"MySqlMaxIdle"`
	MySqlMaxOpen      int         `yaml:"MySqlMaxOpen"`
	SlaveMySqlUrl     string      `yaml:"SlaveMySqlUrl"`
	SlaveMySqlMaxIdle int         `yaml:"SlaveMySqlMaxIdle"`
	SlaveMySqlMaxOpen int         `yaml:"SlaveMySqlMaxOpen"`
	RedisCache        RedisConfig `yaml:"RedisCache"`
	BaseURL           string      `yaml:"BaseURL"`
}

type RedisConfig struct {
	Host      []string `yaml:"Host"`
	Password  string   `yaml:"Password"`
	DB        int      `yaml:"DB"`
	MaxIdle   int      `yaml:"MaxIdle"`
	MaxActive int      `yaml:"MaxActive"`
}

var Instance Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&Instance)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
