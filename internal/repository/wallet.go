package repository

import (
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/internal/models"
	"gorm.io/gorm"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (r *WalletRepo) Create(user *models.User) (*models.Wallet, error) {
	wallet := models.Wallet{
		User: user,
	}
	err := r.db.Model(&wallet).Association("Cards").Append(
		[]models.Card{
			{Name: "My card - UAH", CurrencyCode: "UAH", Balance: 0, Closed: false},
		})
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Get wallet by telegramID
func (r *WalletRepo) Get(telegramID int64) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Joins("User").First(&wallet, "users.telegram_id = ?", telegramID).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

// GetByPublicID get wallet by public ID
func (r *WalletRepo) GetByPublicID(publicID uuid.UUID) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Joins("User").First(&wallet, "users.public_user_id = ?", publicID).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

// CreateCard ...
func (r *WalletRepo) CreateCard(cardName, currencyCode string) (*models.Card, error) {
	card := models.Card{
		Name:         cardName,
		CurrencyCode: currencyCode,
		Balance:      0,
		Closed:       false,
	}
	if err := r.db.FirstOrCreate(&card).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *WalletRepo) UpdateCardBalance(cardID uint64, amount int) (*models.Card, error) {
	r.db.Model(&models.Card{}).Where("id = ?", cardID).Update("amount", amount)
	return nil, nil
}
