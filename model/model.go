package model

type Config struct {
	DbType       string
	DbAddr       string
	DbPort       string
	DbName       string
	DbUser       string
	DbPass       string
	DbSSLmode    string
	MaxOpenConns int
}
