package main

import (
	"net/http"

	"github.com/harshadmanglani/polaris"
	"github.com/labstack/echo/v4"
)

type ExpenseWorkflow struct {
}

func (ew *ExpenseWorkflow) GetWorkflowMeta() polaris.WorkflowMeta {
	return polaris.WorkflowMeta{
		TargetData: nil,
		Builders:   []polaris.IBuilder{},
	}
}

type ValidateExpense struct {
}

type PersistExpense struct {
}

type SettleBalance struct {
}

type ExpenseSettled struct {
}

func createExpense(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{ctx.Get("claims")})
}
