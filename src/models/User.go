package models

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zGuiOs/poupeme-server/src/security"
)

// User model
type User struct {
	ID         uint64    `json:"id,omitempty"`
	UUID       string    `json:"uuid,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// Prepare validate and format user
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if step == "register" {
		user.UUID = uuid.New().String()
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

// validate verify if user is valid
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Nome não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("Email não pode estar em branco")
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("Email inválido")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Senha não pode estar em branco")
	}

	return nil
}

// format remove blank spaces in the corners
func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
