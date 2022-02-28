package comment

import (
	"FinalProjectGolangH8/domain"
	"fmt"
	"time"
)

type Service interface {
	CreateComment(input domain.InputCommentCreate, userId int) (domain.Comment, error)
	GetAll() ([]domain.Comment, error)
	PutPhoto(input domain.UpdateComment, idComment int) (domain.Comment, error)
	GetCommentByID(id int) (domain.Comment, error)
	DeleteComment(id int) (domain.Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateComment(input domain.InputCommentCreate, userId int) (domain.Comment, error) {
	comment := domain.Comment{}
	comment.User_id = userId
	comment.Message = input.Message
	comment.Photo_id = input.PhotoID
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	newComment, err := s.repository.Save(comment)
	if err != nil {
		return newComment, err
	}
	return newComment, nil
}

func (s *service) GetAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	comments, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *service) GetCommentByID(id int) (domain.Comment, error) {
	var comment domain.Comment
	comment, err := s.repository.FindByID(id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (s *service) PutPhoto(input domain.UpdateComment, idComment int) (domain.Comment, error) {
	comment, err := s.repository.FindByID(idComment)
	if err != nil {
		fmt.Println(err)
		return comment, err
	}
	comment.Message = input.Message
	comment.UpdatedAt = time.Now()
	commentResult, err := s.repository.PutComment(comment)
	if err != nil {
		return comment, err
	}
	return commentResult, nil
}

func (s *service) DeleteComment(id int) (domain.Comment, error) {
	comment, err := s.repository.FindByID(id)
	if err != nil {
		return comment, err
	}
	commentResult, err := s.repository.DeleteComment(id)
	if err != nil {
		return comment, err
	}
	return commentResult, nil
}
