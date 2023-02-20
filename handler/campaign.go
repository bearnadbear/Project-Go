package handler

import (
	"fmt"
	"net/http"
	"project/helper"
	sourceCampaign "project/source_campaign"
	sourceUser "project/source_user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService sourceCampaign.Service
}

func NewCampaignHandler(campaignService sourceCampaign.Service) *CampaignHandler {
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

	response := helper.APIResponse("List of campaign", http.StatusOK, "success", sourceCampaign.FormatCampaigns(campaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	var input sourceCampaign.GetCampaignDetailInput

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

	formatter := sourceCampaign.FormatCampaignDetail(campaign)

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var input sourceCampaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(sourceUser.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", sourceCampaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID sourceCampaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData sourceCampaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}

		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(sourceUser.User)
	inputData.User = currentUser

	updateCampaign, err := h.campaignService.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update campaign", http.StatusOK, "success", sourceCampaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UploadImage(c *gin.Context) {
	var input sourceCampaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}

		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(sourceUser.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload camapaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Campaign Image successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
