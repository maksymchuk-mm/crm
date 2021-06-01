package models

import (
	"github.com/pborman/uuid"
	"time"
)

type User struct {
	ID           uint64    `gorm:"primarykey;index" json:"-"`
	PublicUserID uuid.UUID `gorm:"not null;index;type:uuid;default:uuid_generate_v4()" json:"publicUserID"`
	TelegramID   int64     `gorm:"uniqueIndex;not null" json:"telegramID"`
	PasswordID   uint64    `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Password     *Password `orm:"foreignKey:PasswordID" json:"-"`
}

type Password struct {
	ID          uint64    `gorm:"primarykey;index" json:"-"`
	ExpiredDate time.Time `gorm:"not null" json:"expiredDate"`
	Password    string    `gorm:"not null" json:"password"`
	IsValid     bool      `gorm:"default:false" json:"isValid"`
}

func (p *Password) Validate() bool {
	return p.IsValid && p.ExpiredDate.Sub(time.Now()).Seconds() > 0
}
