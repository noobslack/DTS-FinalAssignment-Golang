package repository

import (
	"mygram-finalprojectdts/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(newUser model.User) (model.User, error) {
	tx := ur.db.Create(&newUser)
	return newUser, tx.Error
}

func (ur *userRepository) GetByUsername(email string) (model.User, error) {
	var user model.User

	tx := ur.db.First(&user, "email = ?", email)

	return user, tx.Error
}

func (ur *userRepository) Update(id uint, newUser model.User) (model.User, error) {
	tx := ur.db.Model(&newUser).Where("id=?", id).Updates(model.User{
		Email:    newUser.Email,
		UserName: newUser.UserName,
		Age:      newUser.Age,
	})

	return newUser, tx.Error
}

func (ur *userRepository) Delete(id uint) error {
	tx := ur.db.Delete(&model.User{}, "id=?", id)
	return tx.Error
}
