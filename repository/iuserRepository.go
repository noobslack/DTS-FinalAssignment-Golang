package repository

import "mygram-finalprojectdts/model"

type IUserRepository interface {
	Create(model.User) (model.User, error)
	GetByUsername(string) (model.User, error)
	Update(id uint, newUser model.User) (model.User, error)
	Delete(id uint) error
}
