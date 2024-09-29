package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email        string    `json:"email" gorm:"unique" validate:"required"`
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type UserResponse struct {
	Id    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
	JWT   string `json:"jwt"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		Id:    u.Id.String(),
		Email: u.Email,
	}
}
