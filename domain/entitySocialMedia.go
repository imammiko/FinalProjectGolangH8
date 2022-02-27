package domain

import "time"

type (
	SocialMedia struct {
		ID               int    `gorm:"primary_key"`
		Name             string `gorm:"type:varchar(100);not null" validate:"required"`
		Social_media_Url string `gorm:"not null;" validate:"required"`
		UserId           int
		CreatedAt        time.Time
		UpdatedAt        time.Time
		User             User
	}
)
