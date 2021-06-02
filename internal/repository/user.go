package repository

import (
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/internal/models"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func getPasswordExpireTime() time.Time {
	return time.Now().Add(time.Minute * 5)
}

func (r *UserRepo) Create(telegramID int64, password string) (*models.User, error) {
	user := &models.User{
		TelegramID: telegramID,
		Password: &models.Password{
			ExpiredDate: getPasswordExpireTime(),
			Password:    password,
			IsValid:     true,
		},
	}
	if err := r.db.FirstOrCreate(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Get user by telegram ID
func (r *UserRepo) Get(telegramID int64) (*models.User, error) {
	var user models.User
	if err := r.db.Joins("Password").First(&user, "users.telegram_id = ?", telegramID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByPublicID get user by public ID
func (r *UserRepo) GetByPublicID(publicID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.Joins("Password").First(&user, "users.public_user_id = ?", publicID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
