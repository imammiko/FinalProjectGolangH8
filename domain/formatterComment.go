package domain

import "time"

type CommentCreate struct {
	ID        int       `json:"id"`
	Message   string    `json:"message" binding:"required"`
	Photo_id  int       `json:"photo_id"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentUserGet struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhotoGet struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   int    `json:"user_id"`
}

type CommentGet struct {
	ID        int             `json:"id"`
	Message   string          `json:"message"`
	PhotoID   int             `json:"photo_id"`
	User_Id   int             `json:"user_id"`
	UpdatedAt time.Time       `json:"updated_at"`
	CreatedAt time.Time       `json:"created_at"`
	User      CommentUserGet  `json:"User"`
	Photo     CommentPhotoGet `json:"Photo"`
}

type CommentPut struct {
	ID       int       `json:"id"`
	Message  string    `json:"message"`
	Photo_id int       `json:"photo_id"`
	User_id  int       `json:"user_id"`
	UpdateAt time.Time `json:"updated_at"`
}

func FormatterCommentOutputCreate(comment Comment) CommentCreate {
	return CommentCreate{
		ID:        comment.ID,
		Message:   comment.Message,
		Photo_id:  comment.Photo_id,
		User_id:   comment.User_id,
		CreatedAt: comment.CreatedAt,
	}
}

func FormatterCommentOutputGet(comment Comment) CommentGet {
	return CommentGet{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.Photo_id,
		User_Id:   comment.User_id,
		UpdatedAt: comment.UpdatedAt,
		CreatedAt: comment.CreatedAt,
		User: CommentUserGet{
			ID:       comment.User.ID,
			Email:    comment.User.Email,
			Username: comment.User.Username,
		},
		Photo: CommentPhotoGet{
			ID:        comment.Photo.ID,
			Title:     comment.Photo.Title,
			Caption:   comment.Photo.Caption,
			Photo_url: comment.Photo.Photo_url,
			User_id:   comment.Photo.User_id,
		},
	}
}

func FormatterCommentOutputPut(comment Comment) CommentPut {
	return CommentPut{
		ID:       comment.ID,
		Message:  comment.Message,
		Photo_id: comment.Photo_id,
		User_id:  comment.User_id,
		UpdateAt: comment.UpdatedAt,
	}
}
