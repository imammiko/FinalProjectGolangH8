package socialmedia

type (
	SocialMedia struct {
		ID               int    `gorm:"primary_key"`
		Name             string `gorm:"type:varchar(100)"`
		Social_media_Url string `gorm:"not null;" validate:"required"`
		UserId           int
	}
)
