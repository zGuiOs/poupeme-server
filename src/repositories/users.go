package repositories

import (
	"database/sql"
	"errors"

	"github.com/zGuiOs/poupeme-server/src/models"
)

type users struct {
	db *sql.DB
}

// NewUserRepository create the user repo
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create create a new user
func (repositorie users) Create(user models.User) error {
	statement, err := repositorie.db.Prepare("INSERT INTO users (uuid, name, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repositorie users) FetchByID(UUID string) (uint64, error) {
	var id uint64

	err := repositorie.db.QueryRow("SELECT id FROM users WHERE uuid = ?", UUID).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, errors.New("Usuário não encontrado")
	}

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateUser update an user in database
func (repositorie users) UpdateUser(UUID string, user models.User) error {
	statement, err := repositorie.db.Prepare("UPDATE users SET name = ?, email = ? WHERE uuid = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, UUID); err != nil {
		return err
	}

	return nil
}

// DeleteUser delete an user in database
func (repositorie users) DeleteUser(UUID string) error {
	statement, err := repositorie.db.Prepare("DELETE FROM users WHERE uuid = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(UUID); err != nil {
		return err
	}

	return nil
}

// FetchByEmail fetch an user by email
func (repositorie users) FetchByEmail(email string) (models.User, error) {
	lines, err := repositorie.db.Query("SELECT uuid, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, nil
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.UUID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
