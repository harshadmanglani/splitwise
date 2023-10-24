package main

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/harshadmanglani/splitwise/models"
	"github.com/labstack/echo/v4"
)

func insertUser(ctx echo.Context) error {
	app := ctx.Get("app").(*App)
	var user models.User
	uu, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	user.Uuid = uu.String()
	fmt.Println(user)
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	fmt.Println(user)
	if err := app.queries.InsertUser.Get(&user.Id,
		user.Uuid,
		user.Username,
		user.Name,
		user.Email,
		user.Phone,
		user.PassHash); err != nil {
		fmt.Printf("Error inserting user: %s", err)
	}

	return ctx.JSON(http.StatusOK, okResp{true})
}
