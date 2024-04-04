package config

import (
	"github.com/zGuiOs/poupeme-server/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDataBase() (*gorm.DB, error){
	logger := GetLogger("mysql")

	// String de conexão ao db;
	connectionString := "root:senha@tcp(localhost:3306)/poupeme"

	// Conectando ao db;
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar ao bando de dados: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{}, &schemas.Transaction{})
	if err != nil {
		logger.Errorf("Erro no AutoMigrate: %v", err)
		return nil, err
	}

	return db, nil
}