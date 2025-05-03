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

// Create a new transaction
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

func (repository transactions) FetchTransactions(userID uint64) ([]models.Transaction, error) {
	rows, err := repository.db.Query("SELECT title, description, amount, date, type_id, category_id, status, is_recurring, method FROM transactions WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction

		err := rows.Scan(&transaction.Title, &transaction.Description, &transaction.Amount, &transaction.Date, &transaction.Type_id, &transaction.Category_id, &transaction.Status, &transaction.Is_recurring, &transaction.Method)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
