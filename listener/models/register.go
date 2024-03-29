package models

import "github.com/google/uuid"

// Incoming payload for new user registration
type APIRegisterModel struct {
	CharacterName   string    `json:"charactername"`
	EmailAddress    string    `json:"emailaddress"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmpassword"`
	StartID         uuid.UUID `json:"startid"`
}
