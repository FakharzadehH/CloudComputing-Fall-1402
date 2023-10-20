package config

type Mailgun struct {
	Domain       string `mapstructure:"domain"`
	ApiKey       string `mapstructure:"api_key"`
	PublicApiKey string `mapstructure:"public_api_key"`
}
