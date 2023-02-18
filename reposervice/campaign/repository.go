package reposervice

import (
	model "project/model/campaign"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Campaign, error)
	FindByUserID(UserID int) ([]model.Campaign, error)
	FindByID(ID int) (model.Campaign, error)
	Save(campaign model.Campaign) (model.Campaign, error)
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

func (r *repository) FindByID(ID int) (model.Campaign, error) {
	var campaign model.Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign model.Campaign) (model.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
