package user

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(u User) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) (User, error)
	FindByID(id int) (User, error)
	FindAll() ([]User, error)
	FindUsername(username string) (User, error)
	DeleteUser(id int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(u User) (User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	user.UpdatedAt = time.Now()
	err := r.db.Save(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindUsername(username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) DeleteUser(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
