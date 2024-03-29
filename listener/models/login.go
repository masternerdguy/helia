package models

// Incoming payload for existing user login
type APILoginModel struct {
	EmailAddress string `json:"emailaddress"`
	Password     string `json:"password"`
}

// Outgoing result of user login attempt
type APILoginResponseModel struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	SessionID string `json:"sid"`
	UID       string `json:"uid"`
}
