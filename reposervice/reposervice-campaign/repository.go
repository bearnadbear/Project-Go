package reposerviceCampaign

import (
	modelCampaign "project/model/campaign"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]modelCampaign.Campaign, error)
	FindByUserID(UserID int) ([]modelCampaign.Campaign, error)
	FindByID(ID int) (modelCampaign.Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]modelCampaign.Campaign, error) {
	var campaigns []modelCampaign.Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (r *repository) FindByUserID(UserID int) ([]modelCampaign.Campaign, error) {
	var campaigns []modelCampaign.Campaign

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID int) (modelCampaign.Campaign, error) {
	var campaign modelCampaign.Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
