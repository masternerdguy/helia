package models

//APILoginModel Incoming payload for existing user login
type APILoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//APILoginResponseModel Outgoing result of user login attempt
type APILoginResponseModel struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	SessionID string `json:"sid"`
	UID       string `json:"uid"`
}
