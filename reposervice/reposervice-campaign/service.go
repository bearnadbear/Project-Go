package reposerviceCampaign

import modelCampaign "project/model/campaign"

type Service interface {
	GetCampaign(userID int) ([]modelCampaign.Campaign, error)
	GetCampaignByID(input modelCampaign.GetCampaignDetailInput) (modelCampaign.Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(userID int) ([]modelCampaign.Campaign, error) {
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

func (s *service) GetCampaignByID(input modelCampaign.GetCampaignDetailInput) (modelCampaign.Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
