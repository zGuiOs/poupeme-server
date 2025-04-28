package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/auth"
	"github.com/zGuiOs/poupeme-server/src/database"
	"github.com/zGuiOs/poupeme-server/src/models"
	"github.com/zGuiOs/poupeme-server/src/repositories"
	"github.com/zGuiOs/poupeme-server/src/responses"
	"github.com/zGuiOs/poupeme-server/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepository(db)
	userDB, err := repositorie.FetchByEmail(user.Email)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userDB.Password, user.Password); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userDB.UUID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
}
