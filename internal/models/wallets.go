package models

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
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
	tag := getCurrencyTag(c.CurrencyCode)
	cur, _ := currency.FromTag(tag)
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(c.Balance/100, number.Scale(scale))
	p := message.NewPrinter(tag)
	return p.Sprintf("%v %#v", currency.Symbol(cur), dec)
}

func getCurrencyTag(s string) language.Tag {
	switch s {
	case "USD":
		return language.MustParse("en")
	case "EUR":
		return language.MustParse("eu")
	default:
		return language.MustParse("uk")
	}
}
