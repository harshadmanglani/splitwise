package models

import null "gopkg.in/volatiletech/null.v6"

type Base struct {
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
	UpdatedAt null.Time `db:"updated_at" json:"updatedAt"`
}

type User struct {
	Base

	UserId   int    `db:"user_id" json:"userId"`
	Username string `db:"username" json:"username"`
	PassHash string `db:"pass_hash" json:"-"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
}

type Group struct {
	Base

	GroupId int    `db:"group_id" json:"groupId"`
	Name    string `db:"name" json:"name"`
}

type UserGroupMapping struct {
	Base

	UserId  int    `db:"user_id" json:"userId"`
	GroupId string `db:"group_id" json:"groupId"`
}

type Expense struct {
	Base

	ExpenseId int    `db:"expense_id" json:"expenseId"`
	Title     string `db:"title" json:"title"`
	Amount    int64  `db:"amount" json:"amount"`
	OwedTo    string `db:"owed_to" json:"owedTo"`
	GroupId   string `db:"group_id" json:"groupId"`
}

type Balance struct {
	Base

	BalanceId int    `db:"balance_id" json:"balanceId"`
	Amount    int64  `db:"amount" json:"amount"`
	OwedBy    string `db:"owed_by" json:"owedBy"`
	ExpenseId string `db:"expense_id" json:"expenseId"`
	GroupId   string `db:"group_id" json:"groupId"`
	Settled   bool   `db:"settled" json:"isSettled"`
}
