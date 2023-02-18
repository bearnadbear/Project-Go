package handler

import (
	"net/http"
	"project/helper"
	model "project/model/campaign"
	reposervice "project/reposervice/campaign"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService reposervice.Service
}

func NewCampaignHandler(campaignService reposervice.Service) *CampaignHandler {
	return &CampaignHandler{campaignService}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.campaignService.GetCampaign(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaign", http.StatusOK, "success", model.FormatCampaigns(campaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	var input model.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := model.FormatCampaignDetail(campaign)

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
