package schemas

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title       string
	Description string
	Date        time.Time
	Amount      decimal.Decimal `gorm:"type:decimal(10,2)"`
}
