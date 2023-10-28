package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func createExpense(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{true})
}
