package sourceTransaction

import sourceUser "project/source_user"

type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User sourceUser.User
}
