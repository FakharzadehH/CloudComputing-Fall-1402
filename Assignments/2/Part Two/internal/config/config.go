package config

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	APP       APP       `mapstructure:"app"`
	Redis     Redis     `mapstructure:"redis"`
	NinjasAPI NinjasAPI `mapstructure:"ninjas_api"`
}

type NinjasAPI struct {
	APIKey  string `mapstructure:"api_key"`
	APIHost string `mapstructure:"api_host"`
}

type APP struct {
	CityName string `mapstructure:"city_name"`
}

type Redis struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Expiry int    `mapstructure:"expiry"`
}

func GetConfig() Config {
	return cfg
}

func NewRedisPool(cfg Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		PoolSize: 5, //max number of connections
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
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
