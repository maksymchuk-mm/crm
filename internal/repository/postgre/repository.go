package postgre

import (
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/internal/models"
	"gorm.io/gorm"
)

type User interface {
	Create(telegramID int64, password string) (*models.User, error)
	Get(telegramID int64) (*models.User, error)
	GetByPublicID(publicID uuid.UUID) (*models.User, error)
}

type Wallet interface {
	Create(user *models.User) (*models.Wallet, error)
	Get(telegramID int64) (*models.Wallet, error)
	GetByPublicID(publicID uuid.UUID) (*models.Wallet, error)
}

type Card interface {
}

type Transactions interface {
}

type Type interface {
}

type Repositories struct {
	User         User
	Wallet       Wallet
	Card         Card
	Transactions Transactions
	Type         Type
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: NewUserRepo(db),
	}
}
