package repositories

import (
	"database/sql"

	"github.com/zGuiOs/poupeme-server/src/models"
)

type transactions struct {
	db *sql.DB
}

// NewTransactionRepository create the transaction repo
func NewTransactionRepository(db *sql.DB) *transactions {
	return &transactions{db}
}

func (repository transactions) Create(transaction models.Transaction) error {
	statement, err := repository.db.Prepare("INSERT INTO transactions (uuid, user_id, title, description, amount, date, type_id, category_id, is_recurring, method) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(transaction.UUID, transaction.USER_ID, transaction.Title, transaction.Description, transaction.Amount, transaction.Date, transaction.Type_id, transaction.Category_id, transaction.Is_recurring, transaction.Method)
	if err != nil {
		return err
	}

	return nil
}
