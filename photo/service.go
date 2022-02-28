package photo

import (
	"FinalProjectGolangH8/domain"
	"fmt"
	"time"
)

type Service interface {
	CreatePhoto(input domain.InputPhotos, userId int) (domain.Photo, error)
	GetAll() ([]domain.Photo, error)
	GetPhotoByID(id int) (domain.Photo, error)
	PutPhoto(input domain.InputPhotos, idPhoto int) (domain.Photo, error)
	DeletePhoto(id int) (domain.Photo, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreatePhoto(input domain.InputPhotos, userId int) (domain.Photo, error) {
	photo := domain.Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.Photo_url = input.Photo_url
	photo.User_id = userId
	newPhoto, err := s.repository.Save(photo)
	if err != nil {
		return newPhoto, err
	}
	return newPhoto, nil
}

func (s *service) GetAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	photos, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func (s *service) GetPhotoByID(id int) (domain.Photo, error) {
	var photo domain.Photo
	photo, err := s.repository.FindByID(id)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *service) PutPhoto(input domain.InputPhotos, idPhoto int) (domain.Photo, error) {
	photo, err := s.repository.FindByID(idPhoto)

	if err != nil {
		fmt.Println(err)
		return photo, err
	}
	photo.Title = input.Title
	if input.Caption != "" {
		photo.Caption = input.Caption
	}
	photo.Photo_url = input.Photo_url
	photo.UpdatedAt = time.Now()
	newPhoto, err := s.repository.PutPhoto(photo)
	if err != nil {
		return newPhoto, err
	}
	return newPhoto, nil
}

func (s *service) DeletePhoto(id int) (domain.Photo, error) {
	photo, err := s.repository.FindByID(id)
	if err != nil {
		return photo, err
	}
	_, err = s.repository.DeletePhoto(id)
	if err != nil {
		return photo, err
	}
	return photo, nil
}
