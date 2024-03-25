package repository

import (
	"mygram-finalprojectdts/model"

	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (pr *socialMediaRepository) Create(newSocMed model.SocialMedia) (model.SocialMedia, error) {
	tx := pr.db.Create(&newSocMed)
	return newSocMed, tx.Error

}

func (pr *socialMediaRepository) GetAll() ([]model.SocialMedia, error) {
	var socialMedia = []model.SocialMedia{}
	tx := pr.db.Preload("User").Find(&socialMedia)

	return socialMedia, tx.Error
}

func (pr *socialMediaRepository) Update(id uint, newSocMed model.SocialMedia) (model.SocialMedia, error) {
	tx := pr.db.Model(&newSocMed).Where("id=?", id).Updates(model.SocialMedia{
		Name:           newSocMed.Name,
		SocialMediaURL: newSocMed.SocialMediaURL,
	})

	return newSocMed, tx.Error
}
func (pr *socialMediaRepository) Delete(id uint) error {
	tx := pr.db.Delete(&model.SocialMedia{}, "id = ?", id)
	return tx.Error
}

func (pr *socialMediaRepository) GetOne(id uint) ([]model.SocialMedia, error) {

	var socialMedias = []model.SocialMedia{}
	tx := pr.db.Find(&socialMedias, "id = ?", id)
	return socialMedias, tx.Error
}
