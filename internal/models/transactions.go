package models

import "time"

const (
	TypeIncomeOperations = 1000
	TypeCostsOperations  = 1001
)

type Transaction struct {
	ID          uint64    `gorm:"primaryKey;index" json:"-"`
	TypeID      uint64    `gorm:"not null;constraint:OnDelete:CASCADE;" json:"-"`
	Type        *Type     `gorm:"foreignKey:TypeID" json:"-"`
	Amount      int64     `gorm:"not null;" json:"amount"`
	UserID      uint64    `gorm:"not null;constraint:OnDelete:CASCADE;" json:"-"`
	User        *User     `gorm:"foreignKey:UserID" json:"-"`
	CardID      uint64    `gorm:"not null;constraint:OnDelete:CASCADE;" json:"transaction"`
	Card        *Card     `gorm:"foreignKey:CardID" json:"-"`
	CreatedAt   time.Time `gorm:"index;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"index;autoUpdateTime" json:"updatedAt"`
	Description string    `json:"description"`
}

type Type struct {
	ID   uint64 `gorm:"primaryKey;index" json:"-"`
	Name string `gorm:"not null;" json:"name"`
	Code int    `gorm:"null;" json:"code"`
}
