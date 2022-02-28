package socialmedia

import (
	"FinalProjectGolangH8/domain"
	"FinalProjectGolangH8/photo"
)

type Service interface {
	CreateSocialMedia(input domain.InputSocialMedia, userId int) (domain.SocialMedia, error)
	GetAll() ([]domain.SocialMedia, []domain.Photo, error)
	GetSocialMediaByID(id int) (domain.SocialMedia, error)
	PutSocialMedia(input domain.InputSocialMedia, idSocialMedia int) (domain.SocialMedia, error)
	DeleteSocialMedia(id int) (domain.SocialMedia, error)
}
type service struct {
	repository       Repository
	photosRepository photo.Repository
}

func NewService(repository Repository, photoRepository photo.Repository) Service {
	return &service{
		repository:       repository,
		photosRepository: photoRepository,
	}
}

func (s *service) CreateSocialMedia(input domain.InputSocialMedia, userId int) (domain.SocialMedia, error) {
	socialMedia := domain.SocialMedia{}
	socialMedia.Name = input.Name
	socialMedia.Social_media_Url = input.SocialMeidaUrl
	socialMedia.UserId = userId
	newSocialMedia, err := s.repository.Save(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}
	return newSocialMedia, nil
}

func (s *service) GetAll() ([]domain.SocialMedia, []domain.Photo, error) {
	var socialMedias []domain.SocialMedia
	var photos []domain.Photo
	socialMedias, err := s.repository.FindAll()
	if err != nil {
		return nil, nil, err
	}
	photos, err = s.photosRepository.FindAll()
	if err != nil {
		return nil, nil, err
	}
	return socialMedias, photos, nil
}

func (s *service) GetSocialMediaByID(id int) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	socialMedia, err := s.repository.FindByID(id)
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (s *service) PutSocialMedia(input domain.InputSocialMedia, idSocialMedia int) (domain.SocialMedia, error) {
	socialMedia, err := s.repository.FindByID(idSocialMedia)
	if err != nil {
		return socialMedia, err
	}
	socialMedia.Name = input.Name
	socialMedia.Social_media_Url = input.SocialMeidaUrl
	newSocialMedia, err := s.repository.PutSocialMedia(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}
	return newSocialMedia, nil
}

func (s *service) DeleteSocialMedia(id int) (domain.SocialMedia, error) {
	_, err := s.repository.DeleteSocialMedia(id)
	if err != nil {
		return domain.SocialMedia{}, err
	}
	return domain.SocialMedia{}, nil
}
