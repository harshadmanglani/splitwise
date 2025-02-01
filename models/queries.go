package models

import "github.com/jmoiron/sqlx"

type Queries struct {
	GetUser                         *sqlx.Stmt `query:"get-user"`
	InsertUser                      *sqlx.Stmt `query:"insert-user"`
	InsertGroup                     *sqlx.Stmt `query:"insert-group"`
	InsertUserInGroup               *sqlx.Stmt `query:"insert-user-group-mappings"`
	GetGroup                        *sqlx.Stmt `query:"get-group"`
	GetGroupsForUser                *sqlx.Stmt `query:"get-groups-for-user"`
	InsertExpense                   *sqlx.Stmt `query:"insert-expense"`
	InsertUserInExpense             *sqlx.Stmt `query:"insert-user-expense-mapping"`
	GetExpense                      *sqlx.Stmt `query:"get-expense"`
	GetExpensesInGroup              *sqlx.Stmt `query:"get-expenses-in-group"`
	GetOwedToExpensesForUserInGroup *sqlx.Stmt `query:"get-owed-by-expenses-for-user-in-group"`
	GetOwedByExpensesForUserInGroup *sqlx.Stmt `query:"get-owed-to-expenses-for-user-in-group"`
}
