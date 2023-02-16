package reposervicecampaign

import "project/model/campaign"

type Service interface {
	GetCampaign(userID int) ([]campaign.Campaign, error)
	GetCampaignByID(input campaign.GetCampaignDetailInput) (campaign.Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(userID int) ([]campaign.Campaign, error) {
	if userID != 0 {
		campaign, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) GetCampaignByID(input campaign.GetCampaignDetailInput) (campaign.Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
