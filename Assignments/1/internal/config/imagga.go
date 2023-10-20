package config

type Imagga struct {
	ApiKey    string `mapstructure:"api_key"`
	ApiSecret string `mapstructure:"api_secret"`
}
