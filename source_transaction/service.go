package sourceTransaction

import (
	"errors"
	sourceCampaign "project/source_campaign"
)

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository sourceCampaign.Repository
}

func NewService(repository Repository, campaignRepository sourceCampaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction, err := s.repository.FindCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
