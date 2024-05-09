package userinterface

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllVotes() echo.HandlerFunc {

	return func(ctx echo.Context) error {

		return ctx.String(http.StatusOK, "Hello world 😎")
	}
}

func RegisterNewVote() echo.HandlerFunc {

	return func(ctx echo.Context) error {

		return ctx.String(http.StatusCreated, "Created successfully 😍")
	}
}
