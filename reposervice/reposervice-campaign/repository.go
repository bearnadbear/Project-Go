package reposervicecampaign

import (
	"project/model/campaign"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]campaign.Campaign, error)
	FindByUserID(UserID int) ([]campaign.Campaign, error)
	FindByID(ID int) (campaign.Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (r *repository) FindByUserID(UserID int) ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID int) (campaign.Campaign, error) {
	var campaign campaign.Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
