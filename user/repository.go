package user

import (
	"FinalProjectGolangH8/domain"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(u domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	FindByID(id int) (domain.User, error)
	FindAll() ([]domain.User, error)
	FindUsername(username string) (domain.User, error)
	DeleteUser(id int) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(u domain.User) (domain.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *repository) Update(user domain.User) (domain.User, error) {
	user.UpdatedAt = time.Now()
	err := r.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *repository) FindByID(id int) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (r *repository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *repository) DeleteUser(id int) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
