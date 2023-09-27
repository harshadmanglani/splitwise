package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type okResp struct {
	Data interface{} `json:"data"`
}

func initHTTPHandlers(e *echo.Echo, app *App) {
	var g *echo.Group

	// TODO: implement JWTs
	g.GET("/api/health", handleHealthCheck)
	g.GET("/api/config", handleGetServerConfig)
}

// handleHealthCheck is a healthcheck endpoint that returns a 200 response.
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, okResp{true})
}

// handleGetServerConfig returns general server config.
func handleGetServerConfig(c echo.Context) error {
	// TODO: implement
	return c.JSON(http.StatusOK, okResp{})
}
