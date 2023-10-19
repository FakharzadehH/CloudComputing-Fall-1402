package config

import "fmt"

type RabbitMQ struct {
	Host      string `mapstructure:"host"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Type      string `mapstructure:"type"`
	Port      string `mapstructure:"port"`
	QueueName string `mapstructure:"queue_name"`
}

func (mq RabbitMQ) GetURI() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		mq.Type, mq.Username, mq.Password, mq.Host, mq.Port, mq.Username)
}
