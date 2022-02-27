package domain

import (
	"time"
)

type (
	Photo struct {
		ID        int    `gorm:"primary_key"`
		Title     string `gorm:"not null" validate:"required"`
		Caption   string
		Photo_url string `gorm:"not null" validate:"required"`
		User_id   int
		Comments  []Comment `gorm:"foreignKey:Photo_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt time.Time
		UpdatedAt time.Time
		User      User
	}
)
