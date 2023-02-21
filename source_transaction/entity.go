package sourceTransaction

import (
	sourceUser "project/source_user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       sourceUser.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
