package model

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
