package repository

import (
	"mygram-finalprojectdts/model"

	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		db: db,
	}
}

func (pr *photoRepository) Create(newPhoto model.Photo) (model.Photo, error) {
	tx := pr.db.Create(&newPhoto)
	return newPhoto, tx.Error

}

func (pr *photoRepository) GetAll() ([]model.Photo, error) {
	var photos = []model.Photo{}
	tx := pr.db.Preload("User").Find(&photos)

	return photos, tx.Error
}

func (pr *photoRepository) Update(id uint, newPhoto model.Photo) (model.Photo, error) {
	tx := pr.db.Model(&newPhoto).Where("id=?", id).Updates(model.Photo{
		Title:    newPhoto.Title,
		Caption:  newPhoto.Caption,
		PhotoUrl: newPhoto.PhotoUrl,
	})

	return newPhoto, tx.Error
}
func (pr *photoRepository) Delete(id uint) error {
	tx := pr.db.Delete(&model.Photo{}, "id = ?", id)
	return tx.Error
}

func (pr *photoRepository) GetOne(id uint) ([]model.Photo, error) {

	var photos = []model.Photo{}
	tx := pr.db.Find(&photos, "id = ?", id)
	return photos, tx.Error
}
