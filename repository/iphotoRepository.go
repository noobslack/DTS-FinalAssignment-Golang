package repository

import "mygram-finalprojectdts/model"

type IPhotoRepository interface {
	Create(model.Photo) (model.Photo, error)
	GetAll() ([]model.Photo, error)
	Update(id uint, newUser model.Photo) (model.Photo, error)
	Delete(id uint) error
	GetOne(id uint) ([]model.Photo, error)
}
