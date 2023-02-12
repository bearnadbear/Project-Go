package reserv_campaign

import (
	"project/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Campaign, error)
	FindByUserID(UserID int) ([]model.Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Campaign, error) {
	var campaigns []model.Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (r *repository) FindByUserID(UserID int) ([]model.Campaign, error) {
	var campaigns []model.Campaign

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
