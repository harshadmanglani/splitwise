package main

import (
	"fmt"
	"net/http"

<<<<<<< HEAD
	"github.com/harshadmanglani/splitwise/jwt"
=======
	"github.com/harshadmanglani/whopays/jwt"
>>>>>>> ca4e8dcb09fdce6af5de352209cec842475b85ed
	"github.com/labstack/echo/v4"
)

type okResp struct {
	Data interface{} `json:"data"`
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		app := ctx.Get("app").(*App)
		jwtg := app.jwtg
		auth := ctx.Request().Header.Get("Authorization")
		claims, err := jwtg.VerifyAndReturnClaims(auth)
		if err != jwt.NO_ERROR {
			return ctx.JSON(http.StatusUnauthorized, okResp{err})
		}
		ctx.Set("claims", claims)
		return next(ctx)
	}
}

func initHTTPHandlers(e *echo.Echo) {
	var api *echo.Group = e.Group("/api")
	api.GET("/health", handleHealthCheck)

	var users *echo.Group = e.Group("/users")
	users.POST("", insertUser)
	users.GET("", getUser)
	users.POST("/login", loginUser)

	var expenses *echo.Group = e.Group("/expenses")
	expenses.Use(jwtMiddleware)
	expenses.POST("", createExpense)

	routes := e.Routes()
	fmt.Println("Registered Routes:")
	for _, route := range routes {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
}

// handleHealthCheck is a healthcheck endpoint that returns a 200 response.
func handleHealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{true})
}
