package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGORMConnection(cfg Config) (*gorm.DB, error) {
	pg := postgres.New(postgres.Config{
		DSN:                  cfg.DB.GetURI(),
		PreferSimpleProtocol: true,
	})

	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logger.Error,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	gormDB, err := gorm.Open(pg, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   false,
			TablePrefix:   "",
		},
		PrepareStmt:     false,
		Logger:          log,
		CreateBatchSize: 100,
	})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
