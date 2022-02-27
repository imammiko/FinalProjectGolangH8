package domain

import "time"

type SocialMediaOutputCreate struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaOutputPut struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaaUSer struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Social_media_Url string `json:"social_media_url"`
}

type SocialMediaOutputGet struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	SocialMediaUrl string           `json:"social_media_url"`
	UserId         int              `json:"user_id"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	User           SocialMediaaUSer `json:"User"`
}

type SocialMediaMessage struct {
	Message string `json:"message"`
}

func FormatterSocialMediaMessage(message string) SocialMediaMessage {
	return SocialMediaMessage{
		Message: message,
	}
}

func FormatterSocialMediaOutputCreate(socialMedia SocialMedia) SocialMediaOutputCreate {
	return SocialMediaOutputCreate{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.Social_media_Url,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
	}
}

func FormatterSocialMediaOutputGet(socialMedia SocialMedia, photo Photo) SocialMediaOutputGet {
	return SocialMediaOutputGet{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.Social_media_Url,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
		UpdatedAt:      socialMedia.UpdatedAt,
		User: SocialMediaaUSer{
			ID:               socialMedia.User.ID,
			Username:         socialMedia.User.Username,
			Social_media_Url: photo.Photo_url,
		},
	}
}

func FormatterSocialMediaOutputPUt(socialMedia SocialMedia) SocialMediaOutputPut {
	return SocialMediaOutputPut{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.Social_media_Url,
		UserId:         socialMedia.UserId,
		UpdatedAt:      socialMedia.UpdatedAt,
	}
}
