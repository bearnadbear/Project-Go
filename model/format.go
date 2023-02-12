package model

// Users
type UserFormat struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormat {
	userFormatter := UserFormat{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return userFormatter
}

// Campaigns
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
