package config

import "fmt"

type DB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

const DBMSPostgres = "postgres"

func (db DB) GetURI() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		DBMSPostgres, db.Username, db.Password, db.Host, db.Port, db.DBName)
}
