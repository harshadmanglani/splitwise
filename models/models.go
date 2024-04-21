package models

import null "gopkg.in/volatiletech/null.v6"

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

	Uuid string `db:"uuid" json:"groupId"`
	Name string `db:"name" json:"name"`
}

type UserGroupMapping struct {
	Base

	UserId  string `db:"user_uuid" json:"userId"`
	GroupId string `db:"group_uuid" json:"groupId"`
}

type Expense struct {
	Base

	ExpenseId string `db:"uuid" json:"expenseId"`
	Title     string `db:"title" json:"title"`
	Amount    int64  `db:"amount" json:"amount"`
	OwedTo    string `db:"owed_to" json:"owedTo"`
	GroupId   string `db:"group_uuid" json:"groupId"`
}

type UserExpenseMapping struct {
	Base

	UserExpenseMappingId string `db:"uuid"`
}
