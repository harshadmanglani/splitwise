package models

import "github.com/jmoiron/sqlx"

type Queries struct {
	GetUser       *sqlx.Stmt `query:"get-user"`
	InsertUser    *sqlx.Stmt `query:"insert-user"`
	InsertExpense *sqlx.Stmt `query:"insert-expense"`
	InsertSplit   *sqlx.Stmt `query:"insert-split"`
}
