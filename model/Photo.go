package model

import (
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Photo struct {
	ID       uint   `json:"photo_id"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	UserID   uint
	User     *User
	Comment  *Comment

	gorm.Model
}

func (u *Photo) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) error {
	validate := validator.New()

	validationFields := struct {
		Title    string `json:"title" validate:"required"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url" validate:"required"`
	}{
		Title:    u.Title,
		Caption:  u.Caption,
		PhotoUrl: u.PhotoUrl,
	}
	if err := validate.Struct(validationFields); err != nil {
		return err
	}

	return nil
}

type PostPhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	UserID   uint   `json:"user_id"`
}

type ResponsePhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" `
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UpdatePhotoResponse
}
