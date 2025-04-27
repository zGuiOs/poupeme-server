package models

import "time"

// User model
type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
