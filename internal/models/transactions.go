package models

import (
	"database/sql/driver"
	"github.com/maksymchuk-mm/crm/pkg/utils"
	"time"
)

type Type struct {
	Val  int64
	Name string
}

func (t Type) Value() (driver.Value, error) {
	return t.Val, nil
}

func (t *Type) Scan(value interface{}) error {
	undefined := Type{Val: -1, Name: "undefined"}
	if value == nil {
		*t = undefined
		return nil
	}
	val, err := driver.Int32.ConvertValue(value)
	if err != nil {
		return err
	}
	switch val {
	case TypeIncomeOperations.Val:
		*t = TypeIncomeOperations
	case TypeCostsOperations.Val:
		*t = TypeCostsOperations
	default:
		*t = undefined
	}
	return nil
}

var (
	TypeIncomeOperations = Type{Val: 0, Name: "income"}
	TypeCostsOperations  = Type{Val: 1, Name: "costs"}
)

type Transaction struct {
	ID          uint64    `gorm:"primaryKey;index" json:"-"`
	Type        Type      `json:"type"`
	Amount      int64     `gorm:"not null;" json:"amount"`
	UserID      uint64    `gorm:"not null;constraint:OnDelete:CASCADE;" json:"-"`
	User        *User     `gorm:"foreignKey:UserID" json:"-"`
	CardID      uint64    `gorm:"not null;constraint:OnDelete:CASCADE;" json:"transaction"`
	Card        *Card     `gorm:"foreignKey:CardID" json:"-"`
	CreatedAt   time.Time `gorm:"index;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"index;autoUpdateTime" json:"updatedAt"`
	Description string    `json:"description"`
}

// TODO: Maybe need table Category for transactions

func (t *Transaction) FormatAmount() string {
	return utils.CurrencyFormat(t.Card.CurrencyCode, t.Amount)
}

func (t *Transaction) NameType() string {
	return t.Type.Name
}
