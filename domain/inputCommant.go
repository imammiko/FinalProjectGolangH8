package domain

type InputCommentCreate struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id"`
}

type UpdateComment struct {
	Message string `json:"message" binding:"required" `
}
