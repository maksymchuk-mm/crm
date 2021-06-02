package models

import (
	"github.com/maksymchuk-mm/crm/pkg/utils"
	"time"
)

type Wallet struct {
	ID     uint64 `gorm:"primaryKey;index" json:"-"`
	UserID uint64 `gorm:"not null;constraint:OnDelete:CASCADE;" json:"-"`
	User   *User  `gorm:"foreignKey:UserID" json:"-"`
	Cards  []Card `gorm:"many2many:wallet_cards;" json:"cards"`
}

type Card struct {
	ID           uint64    `gorm:"primaryKey;index" json:"-"`
	Name         string    `gorm:"not null" json:"name"`
	CurrencyCode string    `gorm:"not null;size:3" json:"currencyCode"`
	Balance      int64     `gorm:"default:0" json:"balance"`
	UpdatedAt    time.Time `gorm:"index;autoUpdateTime" json:"updatedAt"`
	Closed       bool      `gorm:"default:false" json:"closed"`
}

func (c *Card) FormatBalance() string {
	return utils.CurrencyFormat(c.CurrencyCode, c.Balance)
}
