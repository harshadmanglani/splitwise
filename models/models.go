package models

import null "gopkg.in/volatiletech/null.v6"

// Base holds common fields shared across models.
type Base struct {
	Id        int       `db:"id" json:"id"`
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
	UpdatedAt null.Time `db:"updated_at" json:"updatedAt"`
}

type User struct {
	Base

	Uuid     string `db:"uuid" json:"uuid"`
	Username string `db:"username" json:"username"`
	PassHash string `db:"pass_hash" json:"passHash"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"first_name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
}
