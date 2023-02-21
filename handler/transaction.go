package handler

import (
	"net/http"
	"project/helper"
	sourceTransaction "project/source_transaction"
	sourceUser "project/source_user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService sourceTransaction.Service
}

func NewTransactionHandler(transactionService sourceTransaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input sourceTransaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(sourceUser.User)
	input.User = currentUser

	transaction, err := h.transactionService.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Error to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfuly to get campaign transaction", http.StatusOK, "success", sourceTransaction.FormatCampaignTransactions(transaction))
	c.JSON(http.StatusOK, response)
}
