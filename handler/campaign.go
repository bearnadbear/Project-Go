package handler

import (
	"net/http"
	"project/model"
	"project/reserv_campaign"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService reserv_campaign.Service
}

func NewCampaignHandler(campaignService reserv_campaign.Service) *CampaignHandler {
	return &CampaignHandler{campaignService}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.campaignService.GetCampaign(userID)
	if err != nil {
		response := model.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.APIResponse("List of campaign", http.StatusOK, "succes", campaign)

	c.JSON(http.StatusOK, response)
}
