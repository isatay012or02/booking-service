package repository

import (
	"booking-service/config"
	"booking-service/internal/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	gormDB *gorm.DB
}

func (d database) CloseDB() (err error) {
	s, err := d.gormDB.DB()
	if err != nil {
		return
	}

	err = s.Close()
	return
}

func NewDB(cfg config.DBSettings) (ports.DB, error) {
	logMode := logger.Error
	if cfg.LogMode {
		logMode = logger.Info
	}

	gormDB, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{
		Logger:      logger.Default.LogMode(logMode),
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	postgresDb, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	postgresDb.SetMaxOpenConns(cfg.MaxOpenConns)
	postgresDb.SetMaxIdleConns(cfg.MaxIdleConns)

	return &database{gormDB: gormDB}, nil
}
