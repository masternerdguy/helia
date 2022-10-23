package models

// Incoming payload for existing user password reset request
type APIResetModel struct {
	EmailAddress string `json:"emailaddress"`
}

// Outgoing result of user password reset request
type APIResetResponseModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
