package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zGuiOs/poupeme-server/src/auth"
	"github.com/zGuiOs/poupeme-server/src/database"
	"github.com/zGuiOs/poupeme-server/src/models"
	"github.com/zGuiOs/poupeme-server/src/repositories"
	"github.com/zGuiOs/poupeme-server/src/responses"
)

// CreateUser call the repo that will create the user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	if err = repositorie.Create(user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, "Usuário criado com sucesso!")
}

// UpdateUser call the repo that will update the user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "UUID")

	userUUID, err := auth.ExtractUUID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userUUID != params {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário diferente do seu"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	if err = repositorie.UpdateUser(params, user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser call the repo that will delete the user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "UUID")

	userUUID, err := auth.ExtractUUID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userUUID != params {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário diferente do seu"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	if err = repositorie.DeleteUser(params); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
