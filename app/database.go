package app

import (
	"github.com/yogs696/skilltest/config"
	"github.com/yogs696/skilltest/pkg/gormadp"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBA *gormadp.DBAdapter

// Start database connection
func dbUp(args *AppArgs) {
	var loglevel logger.LogLevel
	if config.Of.App.Debug() {
		loglevel = logger.Info
	} else {
		loglevel = logger.Silent
	}

	pkgOptions := &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	}

	cfg := gormadp.Config{
		Host:     config.Of.Database.Host,
		Port:     config.Of.Database.Port,
		User:     config.Of.Database.User,
		Password: config.Of.Database.Password,
		DBName:   config.Of.Database.Name,
		Dialect:  gormadp.Dialect(config.Of.Database.Dialect),
		Options:  pkgOptions,
	}
	opts := cfg.Dialect.PgOptions(gormadp.PgConfig{
		SSLMode:  false,
		TimeZone: config.Of.App.TimeZone,
	})

	dba := gormadp.Open(cfg, opts)

	DBA = dba
	printOutUp("New DB connection successfully open")
}

// Stop database connection
func dbDown() {
	printOutDown("Closing current DB connection...")
	DBA.Close()
}
