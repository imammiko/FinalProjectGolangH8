package domain

import "time"

type PhotoOutputCreate struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoOutputGet struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	Photo_url string     `json:"photo_url"`
	User_id   int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	User      UserFormat `json:"user"`
}

type PhotoOutputGets struct {
	PhotoOutputGets []PhotoOutputGet
}

type UserFormat struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type MessagePhoto struct {
	Message string `json:"message"`
}

func FormatPhotoOutputCreate(photo Photo) PhotoOutputCreate {

	return PhotoOutputCreate{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_url: photo.Photo_url,
		User_id:   photo.User_id,
		CreatedAt: photo.CreatedAt,
	}
}

func FormatPhotoOutputGet(photo Photo) PhotoOutputGet {
	return PhotoOutputGet{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_url: photo.Photo_url,
		User_id:   photo.User_id,
		CreatedAt: photo.CreatedAt,
		User: UserFormat{
			Email:    photo.User.Email,
			Username: photo.User.Username,
		},
	}
}

func UserFormatPhoto(user User) UserFormat {
	return UserFormat{
		Email:    user.Email,
		Username: user.Username,
	}
}
