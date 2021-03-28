package models

// Incoming payload for new user registration
type APIRegisterModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
