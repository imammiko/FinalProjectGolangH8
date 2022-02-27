package domain

type InputSocialMedia struct {
	Name           string `json:"name" binding:"required"`
	SocialMeidaUrl string `json:"social_media_url" binding:"required"`
}
