package user

import (
	"FinalProjectGolangH8/comment"
	"FinalProjectGolangH8/photo"
	socialmedia "FinalProjectGolangH8/socialMedia"
	"time"
)

type (
	User struct {
		ID       int    `gorm:"primary_key"`
		Username string `gorm:"not null;unique" validate:"required"`
		Email    string `gorm:"not null;unique" validate:"required,email"`
		Password string `gorm:"not null" validate:"required,email,min=6"`
		Age      int    `gorm:"not null" validate:"required,email,gte=8"`

		Comments    []comment.Comment         `gorm:"foreignKey:User_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		SocialMedia []socialmedia.SocialMedia `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Photos      []photo.Photo             `gorm:"foreignKey:User_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
