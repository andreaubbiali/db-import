package model

type Config struct {
	Database dbConfig
	Files    []filesConfig
}

type dbConfig struct {
	DbType       string
	DbAddr       string
	DbPort       string
	DbName       string
	DbUser       string
	DbPass       string
	DbSSLmode    string
	MaxOpenConns int
}

type filesConfig struct {
	FilePath  string `validate:"required,file"`
	Separator rune   `validate:"required"`
	Fields    map[string]string
}
