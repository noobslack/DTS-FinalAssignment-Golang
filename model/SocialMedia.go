package model

import (
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
	UserID         uint   `json:"user_id"`
	User           *User
	gorm.Model
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

func (u *SocialMedia) BeforeUpdate(tx *gorm.DB) error {
	validate := validator.New()

	validationFields := struct {
		Name           string `json:"name" validate:"required"`
		SocialMediaURL string `json:"social_media_url" validate:"required"`
	}{
		Name:           u.Name,
		SocialMediaURL: u.SocialMediaURL,
	}
	if err := validate.Struct(validationFields); err != nil {
		return err
	}

	return nil
}

type SocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UpdateSocMedResponse
}

type GetOneResponse struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}
