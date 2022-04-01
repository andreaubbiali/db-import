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

	conf := config.Config.Database

	if db.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + conf.DbAddr +
			" port=" + conf.DbPort +
			" dbname=" + conf.DbName +
			" user=" + conf.DbUser +
			" password=" + conf.DbPass +
			" sslmode=" + conf.DbSSLmode,
	}), &gormConfig); err != nil { // disable logger = logger.silence
		log.Fatal("Error in db connection", err)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB, _ := db.DB.DB()
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	log.Println("Connected to database ", conf.DbName)

	return &db
}
