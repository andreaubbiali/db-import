package config

import "aubbiali/db-import/model"

var Config = model.Config{
	DbType:       "",
	DbAddr:       "",
	DbPort:       "",
	DbName:       "",
	DbUser:       "",
	DbPass:       "",
	DbSSLmode:    "",
	MaxOpenConns: 0,
}
