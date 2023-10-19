package config

import (
	"strings"

	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	DB       DB       `mapstructure:"db"`
	S3       S3       `mapstructure:"s3"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq"`
}

func GetConfig() Config {
	return cfg
}
func Load(configPath string) error {
	v := viper.New()
	v.SetEnvPrefix("CloudComputing")
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}
