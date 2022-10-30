package models

import "github.com/google/uuid"

// Incoming payload for existing user password reset request
type APIResetModel struct {
	Token           uuid.UUID `json:"token"`
	UserID          uuid.UUID `json:"userId"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmPassword"`
}

// Outgoing result of user password reset request
type APIResetResponseModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
