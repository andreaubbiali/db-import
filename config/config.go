package config

import "aubbiali/db-import/model"

var Config = model.Config{
	DbType:       "postgres",
	DbAddr:       "localhost",
	DbPort:       "5432",
	DbName:       "strd",
	DbUser:       "strd",
	DbPass:       "strd",
	DbSSLmode:    "disable",
	MaxOpenConns: 100,
}
