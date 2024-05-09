package userinterface

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterClient() echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "registered")
	}
}

func LoginClient() echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "login")
	}
}
