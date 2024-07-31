package main

import (
	"fmt"
	"net/http"

	"github.com/harshadmanglani/splitwise/jwt"
	"github.com/labstack/echo/v4"
)

type okResp struct {
	Data interface{} `json:"data"`
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		app := ctx.Get("app").(*App)
		auth := ctx.Request().Header.Get("Authorization")
		claims, err := app.jwt.VerifyAndReturnClaims(auth)
		if err != jwt.NO_ERROR {
			return ctx.JSON(http.StatusUnauthorized, okResp{err})
		}
		ctx.Set("claims", claims)
		return next(ctx)
	}
}

func initHTTPHandlers(e *echo.Echo) {
	e.GET("/healthcheck", handleHealthCheck)
	e.POST("/api/users", createUser)
	e.POST("/api/users/login", loginUser)

	api := e.Group("/api")
	api.Use(authMiddleware)
	api.GET("/users/:userId", getUser)
	api.POST("/expenses", createExpense)
	api.GET("/expenses/:expenseId", getExpense)
	api.PATCH("/expenses/:expenseId", editExpense)
	api.PATCH("/balances/:balanceId", editBalance)
	api.GET("/balances/:balanceId", getBalance)

	routes := e.Routes()
	fmt.Println("Registered Routes:")
	for _, route := range routes {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
}

func handleHealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{true})
}
