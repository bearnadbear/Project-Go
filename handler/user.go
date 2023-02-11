package handler

import (
	"fmt"
	"net/http"
	"project/model"
	"project/reserv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService reserv.Service
}

func NewUserHandler(userService reserv.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input model.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}

		response := model.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegistrasiUser(input)
	if err != nil {
		response := model.APIResponse("Register account failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := model.FormatterUser(user, "cobatoken")

	response := model.APIResponse("Account has been register", http.StatusOK, "Succes", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input model.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}

		response := model.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		errorsMessage := gin.H{"error": err.Error()}

		response := model.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := model.FormatterUser(user, "cobatoken")

	response := model.APIResponse("Successfuly login", http.StatusOK, "Succes", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input model.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}

		response := model.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsAvailableEmail(input)
	if err != nil {
		errorsMessage := gin.H{"error": "Server error"}

		response := model.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_abailable": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email has been register"
	} else {
		metaMessage = "Email is available"
	}

	response := model.APIResponse(metaMessage, http.StatusOK, "Succes", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvataric(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := model.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "Error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := model.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "Error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := model.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "Error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := model.APIResponse("Avatar successfuly uploaded", http.StatusOK, "Succes", data)

	c.JSON(http.StatusOK, response)
}
