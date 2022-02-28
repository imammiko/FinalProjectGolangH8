package socialmedia

import (
	"FinalProjectGolangH8/domain"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Save(u domain.SocialMedia) (domain.SocialMedia, error)
	FindAll() ([]domain.SocialMedia, error)
	PutSocialMedia(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	FindByID(id int) (domain.SocialMedia, error)
	DeleteSocialMedia(id int) (domain.SocialMedia, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(u domain.SocialMedia) (domain.SocialMedia, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) FindAll() ([]domain.SocialMedia, error) {
	var socialMedias []domain.SocialMedia
	err := r.db.Preload(clause.Associations).Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(socialMedias)
	return socialMedias, nil
}

func (r *repository) FindByID(id int) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Where("id = ?", id).First(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *repository) PutSocialMedia(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Save(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *repository) DeleteSocialMedia(id int) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Where("id = ?", id).Delete(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}
