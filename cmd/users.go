package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/harshadmanglani/whopays/jwt"
	"github.com/harshadmanglani/whopays/models"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
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
		if pqErr, ok := err.(*pq.Error); ok {
			// TODO: move to constants
			errMsg := ""
			switch pqErr.Constraint {
			case "users_username_key":
				errMsg = "USERNAME_ALREADY_EXISTS"
			case "users_email_key":
				errMsg = "EMAIL_ALREADY_EXISTS"
			case "users_phone_key":
				errMsg = "PHONE_ALREADY_EXISTS"
			default:
				errMsg = pqErr.Message
			}
			return echo.NewHTTPError(http.StatusConflict, errMsg)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusCreated, okResp{true})
}

func getUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, okResp{models.User{}})
}

func loginUser(ctx echo.Context) error {
	app := ctx.Get("app").(*App)
	jwtg := app.jwtg
	expirationDate := time.Now().Add(2 * time.Minute)

	var request models.LoginRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	var user models.User
	var err error
	if user, err = verifyPassword(app, request); err != nil {
		return err
	}

	claims := jwt.Claims{
		Issuer:   "backend",
		Expiry:   expirationDate,
		Subject:  user.Uuid,
		IssuedAt: time.Now(),
	}
	token := jwtg.GenerateJwt(claims)
	return ctx.JSON(http.StatusOK, okResp{models.LoginResponse{
		User:        user,
		AccessToken: token,
	}})
}

func verifyPassword(app *App, request models.LoginRequest) (models.User, error) {
	var user models.User
	if err := app.queries.GetUser.Get(&user, request.Username, "", ""); err != nil {
		fmt.Printf("User not found, err: %s", err)
		return models.User{}, echo.NewHTTPError(http.StatusNotFound, "INVALID_USER_OR_PASSWORD")
	}

	if request.PassHash != user.PassHash {
		fmt.Printf("Incorrect passHash")
		return models.User{}, echo.NewHTTPError(http.StatusNotFound, "INVALID_USER_OR_PASSWORD")
	}

	return user, nil
}
