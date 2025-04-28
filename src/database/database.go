package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver mysql
	"github.com/zGuiOs/poupeme-server/src/config"
)

// Connect open connection with database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
