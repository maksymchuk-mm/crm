package repository

import (
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/internal/models"
	"github.com/maksymchuk-mm/crm/pkg/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

const limitLastTransaction = 10

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

// Create by slice transactions
func (r *TransactionRepo) Create(trns []models.Transaction) ([]models.Transaction, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, trn := range trns {
			if err := tx.Create(&trn).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return trns, nil
}

// GetUserTransactions returned user transactions by pagination
func (r *TransactionRepo) GetUserTransactions(publicID uuid.UUID, pagination *schemas.Pagination) (*schemas.Pagination, error) {
	var trns []*models.Transaction
	totalRows, totalPages, fromRow, toRow := int64(0), 0, 0, 0
	offset := 0
	if pagination.Page >= 1 {
		offset = (pagination.Page - 1) * pagination.Limit
	}
	err := r.db.Preload(clause.Associations).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where(
		&models.Transaction{
			User: &models.User{
				PublicUserID: publicID,
			},
		}).Find(&trns).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = trns
	err = r.db.Model(&models.Transaction{}).Count(&totalRows).Error
	if err != nil {
		return nil, err
	}
	pagination.TotalRows = totalRows
	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	if pagination.Page == 1 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = (pagination.Page-1)*pagination.Limit + 1
			toRow = pagination.Page * pagination.Limit
		}
	}
	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil
}

// GetUserLastTransactions returned last 10 user transaction
func (r *TransactionRepo) GetUserLastTransactions(publicID uuid.UUID) ([]models.Transaction, error) {
	var trns []models.Transaction
	err := r.db.Preload(clause.Associations).Limit(limitLastTransaction).Order("id desc").Where(
		&models.Transaction{
			User: &models.User{
				PublicUserID: publicID,
			},
		}).Find(&trns).Error
	if err != nil {
		return nil, err
	}
	return trns, nil
}
