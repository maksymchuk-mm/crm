package repository

import (
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/internal/models"
	"github.com/maksymchuk-mm/crm/pkg/schemas"
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

type Transactions interface {
	Create(trns []models.Transaction) ([]models.Transaction, error)
	GetUserTransactions(publicID uuid.UUID, pagination *schemas.Pagination) (*schemas.Pagination, error)
	GetUserLastTransactions(publicID uuid.UUID) ([]models.Transaction, error)
}

type Repositories struct {
	User         User
	Wallet       Wallet
	Transactions Transactions
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:   NewUserRepo(db),
		Wallet: NewWalletRepo(db),
	}
}
