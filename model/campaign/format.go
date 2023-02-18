package model

import "strings"

type CampaignFormat struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	Slug             string `json:"slug"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormat {
	campaignFormatter := CampaignFormat{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		Slug:             campaign.Slug,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormat {
	campaignsFormatter := []CampaignFormat{}

	for _, c := range campaigns {
		campaignFormat := FormatCampaign(c)
		campaignsFormatter = append(campaignsFormatter, campaignFormat)
	}

	return campaignsFormatter
}

type CampaignDetailFormat struct {
	ID               int                   `json:"id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	ImageURL         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserID           int                   `json:"user_id"`
	Slug             string                `json:"slug"`
	Perks            []string              `json:"perks"`
	User             CampaignUserFormat    `json:"user"`
	Images           []CampaignImageFormat `json:"images"`
}

type CampaignUserFormat struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormat struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormat {
	campaignDetailFormatter := CampaignDetailFormat{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	// Next step
	campaignUser := CampaignUserFormat{}
	user := campaign.User

	campaignUser.Name = user.Name
	campaignUser.ImageURL = user.AvatarFileName

	campaignDetailFormatter.Perks = perks
	campaignDetailFormatter.User = campaignUser

	// End step
	images := []CampaignImageFormat{}

	for _, i := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormat{}
		campaignImageFormatter.ImageURL = i.FileName

		isPrimary := false

		if i.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
