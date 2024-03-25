package model

import (
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" validate:"required"`
	User    *User
	Photo   *Photo
	gorm.Model
}

func (u *Comment) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

func (u *Comment) BeforeUpdate(tx *gorm.DB) error {
	validate := validator.New()

	validationFields := struct {
		Message string `json:"message" validate:"required"`
	}{
		Message: u.Message,
	}

	if err := validate.Struct(validationFields); err != nil {
		return err
	}

	return nil
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      UpdateCommentResponse
	Photo     PostPhoto
}
