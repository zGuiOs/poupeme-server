package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;references:UserID"`
  Name     	string `gorm:"size:255"`
  Email    	string `gorm:"unique"`
  Password 	string
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}