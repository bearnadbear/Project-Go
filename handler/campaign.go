package handler

import (
	"net/http"
	"project/model"
	camp "project/model/campaign"
	reposerviceCampaign "project/reposervice/reposervice-campaign"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService reposerviceCampaign.Service
}

func NewCampaignHandler(campaignService reposerviceCampaign.Service) *CampaignHandler {
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

	response := model.APIResponse("List of campaign", http.StatusOK, "success", camp.FormatCampaigns(campaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	var input camp.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := model.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		response := model.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := camp.FormatCampaignDetail(campaign)

	response := model.APIResponse("Campaign detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
