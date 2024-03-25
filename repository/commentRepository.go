package repository

import (
	"mygram-finalprojectdts/model"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (pr *commentRepository) Create(newComment model.Comment) (model.Comment, error) {

	var photos model.Photo
	if err := pr.db.First(&photos, newComment.PhotoID).Error; err != nil {
		return model.Comment{}, err
	}

	err := pr.db.Create(&newComment).Error
	return newComment, err
}

func (pr *commentRepository) GetAll() ([]model.Comment, error) {
	var comments = []model.Comment{}
	tx := pr.db.Preload("User").Preload("Photo").Find(&comments)

	return comments, tx.Error
}

func (pr *commentRepository) Update(id uint, newComment model.Comment) (model.Comment, error) {
	tx := pr.db.Model(&newComment).Where("id=?", id).Updates(model.Comment{
		Message: newComment.Message,
	})

	return newComment, tx.Error
}
func (pr *commentRepository) Delete(id uint) error {
	tx := pr.db.Delete(&model.Comment{}, "id = ?", id)
	return tx.Error
}

func (pr *commentRepository) GetOne(id uint) ([]model.Comment, error) {

	var comments = []model.Comment{}
	tx := pr.db.Find(&comments, "id = ?", id)
	return comments, tx.Error
}
