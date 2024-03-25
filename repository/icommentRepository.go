package repository

import "mygram-finalprojectdts/model"

type ICommentRepository interface {
	Create(model.Comment) (model.Comment, error)
	GetAll() ([]model.Comment, error)
	Update(id uint, newUser model.Comment) (model.Comment, error)
	Delete(id uint) error
	GetOne(id uint) ([]model.Comment, error)
}
