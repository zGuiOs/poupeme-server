package schemas

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId      uint `gorm:"foreignKey:UserID"`
	Title       string
	Description string
	Date        time.Time
	Amount      decimal.Decimal `gorm:"type:decimal(10,2)"`
	Type        string `gorm:"type:enum('revenue', 'expense')"`
	Category    string
	CreatedAt   time.Time `gorm:"created_at"`
  UpdatedAt   time.Time `gorm:"updated_at"`
}
