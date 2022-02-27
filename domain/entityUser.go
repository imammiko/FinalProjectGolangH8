package domain

import (
	"time"
)

type (
	User struct {
		ID       int    `gorm:"primary_key"`
		Username string `gorm:"not null;unique" validate:"required"`
		Email    string `gorm:"not null;unique" validate:"required,email"`
		Password string `gorm:"not null" validate:"required,email,min=6"`
		Age      int    `gorm:"not null" validate:"required,email,gte=8"`

		Comments    []Comment     `gorm:"foreignKey:User_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		SocialMedia []SocialMedia `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Photos      []Photo       `gorm:"foreignKey:User_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
