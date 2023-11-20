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

	UserId   string `db:"uuid" json:"userId"`
	Username string `db:"username" json:"username"`
	PassHash string `db:"pass_hash" json:"-"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
}

type Group struct {
	Base

	GroupId  string `db:"uuid" json:"groupId"`
	Title    string `db:"title" json:"title"`
	Simplify bool   `db:"simplify" json:"isSimplifyDebts"`
}

type Expense struct {
	Base

	ExpenseId   string `db:"uuid" json:"expenseId"`
	Amount      int    `db:"amount" json:"amount"`
	Title       string `db:"title" json:"title"`
	SplitMode   string `db:"split_mode" json:"splitMode"`
	OwnerUserId string `db:"owner_uuid" json:"ownerUserId"`
	GroupId     string `db:"group_uuid" json:"groupId"`
}

// It makes sense to store splits in a separate table as it is graphical data.
// If `users` table holds the nodes, this table holds the edges
type Split struct {
	Base

	SplitId      string `db:"uuid" json:"splitId"`
	ExpenseId    string `db:"expense_uuid" json:"expenseId"`
	GroupId      string `db:"group_uuid" json:"groupId"`
	Amount       int    `db:"amount" json:"amount"`
	OwedByUserId string `db:"owed_by_uuid" json:"owedByUserId"`
	OwedToUserId string `db:"owed_to_uuid" json:"owedToUserId"`
	Settled      bool   `db:"settled" json:"isSettled"`
}
