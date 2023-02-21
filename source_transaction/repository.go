package sourceTransaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindCampaignID(campaingID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCampaignID(campaingID int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaingID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
