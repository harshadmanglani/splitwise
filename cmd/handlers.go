package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type okResp struct {
	Data interface{} `json:"data"`
}

func initHTTPHandlers(e *echo.Echo, app *App) {
	var api *echo.Group = e.Group("/api")
	api.GET("/health", handleHealthCheck)

	var users *echo.Group = e.Group("/users")
	users.POST("", insertUser)

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
