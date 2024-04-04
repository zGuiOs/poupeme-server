package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db 			*gorm.DB
	logger 	*Logger
)

func Init() error {
	var err error

	// Iniciando db;
	db, err = InitializeDataBase()
	if err != nil {
		return fmt.Errorf("Erro iniciando o banco de dados: %v", err)
	}

	return nil
}

func GetMysql() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Iniciando logger
	logger := NewLogger(prefix)
	return logger
}