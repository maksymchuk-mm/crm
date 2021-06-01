package models

import "time"

type Wallet struct {
	ID     uint64 `gorm:"primaryKey;index" json:"-"`
	UserID uint64 `gorm:"not null;constraint:OnDelete:CASCADE;" json:"-"`
	User   *User  `gorm:"foreignKey:UserID" json:"-"`
	Cards  []Card `gorm:"many2many:wallet_cards;" json:"cards"`
}

type Card struct {
	ID         uint64    `gorm:"primaryKey;index" json:"-"`
	Name       string    `gorm:"not null" json:"name"`
	CurrencyID uint      `gorm:"not null;constraint:OnDelete:CASCADE;" json:"-"`
	Currency   *Currency `gorm:"foreignKey:UserID" json:"-"`
	Balance    int64     `gorm:"default:0" json:"balance"`
	UpdatedAt  time.Time `gorm:"index;autoUpdateTime" json:"updatedAt"`
	Closed     bool      `gorm:"default:false" json:"closed"`
}

type Currency struct {
	ID   uint   `gorm:"primaryKey;index" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
