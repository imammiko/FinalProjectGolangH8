package domain

type InputPhotos struct {
	Title     string `json:"title" binding:"required"  `
	Caption   string `json:"caption" `
	Photo_url string `json:"photo_url" binding:"required"`
}
