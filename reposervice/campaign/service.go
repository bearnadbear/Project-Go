package reposervice

import (
	"fmt"
	model "project/model/campaign"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaign(userID int) ([]model.Campaign, error)
	GetCampaignByID(input model.GetCampaignDetailInput) (model.Campaign, error)
	CreateCampaign(input model.CreateCampaignInput) (model.Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(userID int) ([]model.Campaign, error) {
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

func (s *service) GetCampaignByID(input model.GetCampaignDetailInput) (model.Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input model.CreateCampaignInput) (model.Campaign, error) {
	campaign := model.Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		UserID:           input.User.ID,
	}

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
