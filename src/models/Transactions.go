package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           uint64    `json:"id,omitempty"`
	UUID         string    `json:"uuid,omitempty"`
	USER_ID      uint64    `json:"user_id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	Amount       uint64    `json:"amount,omitempty"`
	Date         time.Time `json:"date"`
	Type_id      uint64    `json:"type_id,omitempty"`
	Category_id  uint64    `json:"category_id,omitempty"`
	Status       string    `json:"status,omitempty"`
	Is_recurring bool      `json:"is_recurring,omitempty"`
	Method       string    `json:"method,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

func (transaction *Transaction) Prepare(step string, dateStr string, amountStr string) error {
	if err := transaction.validate(step); err != nil {
		return err
	}

	if step == "register" {
		transaction.UUID = uuid.New().String()
	}

	if err := transaction.format(); err != nil {
		return err
	}

	return nil
}

func (transaction *Transaction) validate(step string) error {
	var validationErrors []string

	if step != "register" && transaction.UUID == "" {
		validationErrors = append(validationErrors, "UUID é obrigatório")
	}

	if transaction.USER_ID == 0 {
		validationErrors = append(validationErrors, "ID do usuário é obrigatório")
	}

	if transaction.Title == "" {
		validationErrors = append(validationErrors, "Título é obrigatório")
	} else if len(transaction.Title) > 50 {
		validationErrors = append(validationErrors, "Título não pode exceder 50 caracteres")
	}

	if transaction.Amount == 0 {
		validationErrors = append(validationErrors, "Valor é obrigatório")
	}

	if transaction.Date.IsZero() {
		validationErrors = append(validationErrors, "Data é obrigatória")
	}

	if transaction.Type_id == 0 {
		validationErrors = append(validationErrors, "ID do tipo é obrigatório")
	}

	if transaction.Category_id == 0 {
		validationErrors = append(validationErrors, "ID da categoria é obrigatório")
	}

	if step != "register" {
		validStatuses := map[string]bool{
			"PENDENTE":  true,
			"REALIZADA": true,
			"CANCELADA": true,
		}
		if transaction.Status == "" {
			validationErrors = append(validationErrors, "Status é obrigatório")
		} else if !validStatuses[strings.ToUpper(transaction.Status)] {
			validationErrors = append(validationErrors, "Status deve ser um de: PENDENTE, REALIZADA, CANCELADA")
		}
	}

	validMethods := map[string]bool{
		"DINHEIRO":      true,
		"DEBITO":        true,
		"CREDITO":       true,
		"PIX":           true,
		"TRANSFERENCIA": true,
	}
	if transaction.Method == "" {
		validationErrors = append(validationErrors, "Método é obrigatório")
	} else if !validMethods[strings.ToUpper(transaction.Method)] {
		validationErrors = append(validationErrors, "Método deve ser um de: DINHEIRO, DEBITO, CREDITO, PIX, TRANSFERENCIA")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}

func (transaction *Transaction) format() error {
	transaction.Title = strings.TrimSpace(transaction.Title)
	transaction.Description = strings.TrimSpace(transaction.Description)

	return nil
}
