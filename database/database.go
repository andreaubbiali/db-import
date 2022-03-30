package database

import (
	"aubbiali/db-import/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	*gorm.DB
}

func SetupDatabase() *database {
	db := database{}

	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false, //if false keeps adding foreign key (duplicates)
		Logger:                                   newLogger,
	}

	if db.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + config.Config.DbAddr +
			" port=" + config.Config.DbPort +
			" dbname=" + config.Config.DbName +
			" user=" + config.Config.DbUser +
			" password=" + config.Config.DbPass +
			" sslmode=" + config.Config.DbSSLmode,
	}), &gormConfig); err != nil { // disable logger = logger.silence
		log.Fatal("Error in db connection", err)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB, _ := db.DB.DB()
	sqlDB.SetMaxOpenConns(config.Config.MaxOpenConns)

	log.Println("Connected to database ", config.Config.DbName)

	return &db
}
