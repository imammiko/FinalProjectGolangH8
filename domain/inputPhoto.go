package domain

type InputPhotos struct {
	Title     string `json:"title" binding:"required"  `
	Caption   string `json:"caption" `
	Photo_url string `json:"photo_url" binding:"required"`
}

type InputPhotosCloud struct {
	Title   string `form:"title"`
	Caption string `form:"caption"`
	// Photo_url string `form:"photo_url" binding:"required"`
}
