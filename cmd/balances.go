package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getBalance(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{ctx.Get("claims")})
}

func editBalance(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{ctx.Get("claims")})
}
