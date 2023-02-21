package sourceTransaction

import "time"

type CampaignTransactionFormat struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	// UpdateAt  time.Time
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormat {
	formatter := CampaignTransactionFormat{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormat {
	if len(transactions) == 0 {
		return []CampaignTransactionFormat{}
	}

	var transactiosFormatter []CampaignTransactionFormat

	for _, v := range transactions {
		formatter := FormatCampaignTransaction(v)
		transactiosFormatter = append(transactiosFormatter, formatter)
	}

	return transactiosFormatter
}
