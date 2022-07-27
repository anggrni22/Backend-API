package io

import (
	"akrab-bangkit2022-api/config"
	"database/sql"
	"fmt"
	l "log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(isLogged bool) *gorm.DB {
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.DBName,
		config.C.Database.Password,
		config.C.Database.SslMode,
	)
	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		l.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		l.Fatalln(err)
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	tm := time.Minute * time.Duration(20)
	sqlDB.SetConnMaxLifetime(tm)

	return db
}

func InitDBS(isLogged bool) *sql.DB {
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{}

	if isLogged {
		newLogger := logger.New(
			l.New(os.Stdout, "\r\n", l.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)

		gormConfig.Logger = newLogger
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.DBName,
		config.C.Database.Password,
		config.C.Database.SslMode,
	)
	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		l.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		l.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	tm := time.Minute * time.Duration(20)
	sqlDB.SetConnMaxLifetime(tm)

	return sqlDB
}
