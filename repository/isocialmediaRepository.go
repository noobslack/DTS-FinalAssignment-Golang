package repository

import "mygram-finalprojectdts/model"

type ISocialMediaRepository interface {
	Create(model.SocialMedia) (model.SocialMedia, error)
	GetAll() ([]model.SocialMedia, error)
	Update(id uint, newSocMed model.SocialMedia) (model.SocialMedia, error)
	Delete(id uint) error
	GetOne(id uint) ([]model.SocialMedia, error)
}
