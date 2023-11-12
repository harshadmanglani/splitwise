package models

type CreateUserResponse struct {
	Uuid string `json:"uuid"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`
}
