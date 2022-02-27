package comment

import (
	"FinalProjectGolangH8/domain"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Save(u domain.Comment) (domain.Comment, error)
	FindAll() ([]domain.Comment, error)
	PutComment(comment domain.Comment) (domain.Comment, error)
	FindByID(id int) (domain.Comment, error)
	DeleteComment(id int) (domain.Comment, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) Save(u domain.Comment) (domain.Comment, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) FindAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Preload(clause.Associations).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *repository) PutComment(comment domain.Comment) (domain.Comment, error) {
	comment.UpdatedAt = time.Now()
	err := r.db.Model(&comment).Updates(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) FindByID(id int) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) DeleteComment(id int) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}
