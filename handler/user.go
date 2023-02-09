package handler

import (
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
		c.JSON(http.StatusBadRequest, response)
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
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		errorsMessage := gin.H{"error": err.Error()}

		response := model.APIResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := model.FormatterUser(user, "cobatoken")

	response := model.APIResponse("Successfuly login", http.StatusOK, "Succes", formatter)

	c.JSON(http.StatusOK, response)
}
