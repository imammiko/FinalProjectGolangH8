package photo

import (
	"FinalProjectGolangH8/domain"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Save(u domain.Photo) (domain.Photo, error)
	FindAll() ([]domain.Photo, error)
	FindByID(id int) (domain.Photo, error)
	PutPhoto(photo domain.Photo) (domain.Photo, error)
	DeletePhoto(id int) (domain.Photo, error)
	FindByUserId(id int) (domain.Photo, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(u domain.Photo) (domain.Photo, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Preload(clause.Associations).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(photos)
	return photos, nil
}

func (r *repository) FindByID(id int) (domain.Photo, error) {
	var photo domain.Photo
	fmt.Println(id)
	err := r.db.Where("id = ?", id).First(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}
func (r *repository) FindByUserId(id int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Where("user_id = ?", id).First(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) PutPhoto(photo domain.Photo) (domain.Photo, error) {
	photo.UpdatedAt = time.Now()
	err := r.db.Save(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) DeletePhoto(id int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}
