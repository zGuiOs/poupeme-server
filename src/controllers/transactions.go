package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zGuiOs/poupeme-server/src/auth"
	"github.com/zGuiOs/poupeme-server/src/database"
	"github.com/zGuiOs/poupeme-server/src/models"
	"github.com/zGuiOs/poupeme-server/src/repositories"
	"github.com/zGuiOs/poupeme-server/src/responses"
)

// CreateTransaction call the repo that will create the transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "UUID")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	userUUID, err := auth.ExtractUUID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userUUID != params {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível criar um transação em um usuário diferente do seu"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userID, err := userRepository.FetchByID(params)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	var transaction models.Transaction
	if err = json.Unmarshal(body, &transaction); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	transaction.USER_ID = userID

	if err = transaction.Prepare("register", transaction.Date.String(), strconv.FormatUint(transaction.Amount, 10)); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	transactionRepository := repositories.NewTransactionRepository(db)
	if err = transactionRepository.Create(transaction); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FetchTransactions call the repo that will fetch all transactions
func FetchTransactions(w http.ResponseWriter, r *http.Request) {}

// FetchTransactions call the repo that will fetch one transaction
func FetchTransactionByID(w http.ResponseWriter, r *http.Request) {}

// UpdateTransaction call the repo that will update the transaction
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {}

// DeleteTransaction call the repo that will delete the transaction
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {}
