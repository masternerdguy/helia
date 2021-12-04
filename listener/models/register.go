package models

// Incoming payload for new user registration
type APIRegisterModel struct {
	CharacterName string `json:"charactername"`
	Password      string `json:"password"`
}
