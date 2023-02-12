package model

import "time"

// Users
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	Password       string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Campaigns
