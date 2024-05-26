package models

type UserLoginRequest struct {
	Email    string`json:"email"`
	Password string`json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string`json:"access_token"`
	RefreshToken string`json:"refresh_token"`
}

type AuthInfo struct {
	UserID string`json:"user_id"`
}