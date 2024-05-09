package main

import (
	"context"
	"fmt"

	// "net/http"

	"voteapp/infrastructure"
	userinterface "voteapp/userInterface"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"os"
)

// func availableMiddleware(ctx context.Context) echo.HandlerFunc {

// 	return func(next echo.HandlerFunc) echo.HandlerFunc  {

// 		isAvailable := ctx.Value("isAvailableApp")

// 		if !isAvailable {
// 			// return e.String(http.StatusServiceUnavailable, "Service unavailable")
// 		}

// 	}
// }

func loadEnvironmentVars() {

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" || appEnv == "development" {
		godotenv.Load("./.env")
	}

}

func main() {

	fmt.Println("app started successfully ...")

	loadEnvironmentVars()

	ctxBackground := context.Background()

	db := infrastructure.InitDB(ctxBackground, os.Getenv("DB_DSN"))

	e := echo.New()

	apiV1Grope := e.Group("/api/v1")

	clients := apiV1Grope.Group("/clients")
	clients.POST("/register", userinterface.RegisterClient(db))
	clients.POST("/login", userinterface.LoginClient(db))

	votes := apiV1Grope.Group("/votes")
	votes.GET("/all-votes", userinterface.GetAllVotes())
	votes.POST("/register-vote", userinterface.RegisterNewVote())

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))

}
