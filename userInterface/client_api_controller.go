package userinterface

import (
	"net/http"
	infra "voteapp/infrastructure"

	"github.com/labstack/echo/v4"
)

func RegisterClient(db infra.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "registered")
	}
}

func LoginClient(db infra.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "login")
	}
}
