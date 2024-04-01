package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db 			*gorm.DB
	logger 	*Logger
)

func Init() error {
	return errors.New("Fake error")
}

func GetLogger(prefix string) *Logger {
	// Iniciando logger
	logger := NewLogger(prefix)
	return logger
}