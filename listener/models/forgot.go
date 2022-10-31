package models

// Incoming payload for existing user password reset request
type APIForgotModel struct {
	EmailAddress string `json:"emailaddress"`
}

// Outgoing result of user password reset request
type APIForgotResponseModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
