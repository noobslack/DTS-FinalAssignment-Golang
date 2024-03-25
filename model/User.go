package model

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type User struct {
	Email        string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	UserName     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required,min=6"`
	Age          uint   `json:"age" validate:"required,gte=8,lte=120"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia

	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	validate := validator.New()

	validationFields := struct {
		UserName string `validate:"required"`
		Email    string `validate:"required,email"`
	}{
		UserName: u.UserName,
		Email:    u.Email,
	}
	if err := validate.Struct(validationFields); err != nil {
		return err
	}

	return nil
}

type UpdatePhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateCommentResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateSocMedResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
