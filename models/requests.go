package models

type CreateUserRequest struct {
	Username string `json:"username"`
	PassHash string `json:"passHash"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Username string `json:"username"`
	PassHash string `json:"passHash"`
}
