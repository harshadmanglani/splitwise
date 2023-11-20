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

type split struct {
	UserUuid string `json:"userId"`
	Value    int    `json:"value"`
}

type CreateExpenseRequest struct {
	Title     string  `json:"title"`
	Amount    int     `json:"amount"`
	SplitMode string  `json:"splitMode"`
	Splits    []split `json:"splits"`
}
