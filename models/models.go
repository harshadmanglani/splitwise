package models

import null "gopkg.in/volatiletech/null.v6"

// Base holds common fields shared across models.
type Base struct {
	id        int       `db:"id" json:"id"`
	createdAt null.Time `db:"created_at" json:"createdAt"`
	updatedAt null.Time `db:"updated_at" json:"updatedAt"`
}

type User struct {
	Base

	uuid      string `db:"uuid" json:"uuid"`
	username  string `db:"username" json:"username"`
	passHash  string `db:"pass_hash" json:"passHash"`
	email     string `db:"email" json:"email"`
	firstName string `db:"first_name" json:"firstName"`
	lastName  string `db:"last_name" json:"lastName"`
	phone     string `db:"phone" json:"phone"`
}
