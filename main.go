package main

import (
	"context"
	"fmt"

	// "net/http"
	"voteapp/infrastructure"
	userinterface "voteapp/userInterface"

	"github.com/labstack/echo/v4"
)

// func availableMiddleware(ctx context.Context) echo.HandlerFunc {

// 	return func(next echo.HandlerFunc) echo.HandlerFunc  {

// 		isAvailable := ctx.Value("isAvailableApp")

// 		if !isAvailable {
// 			// return e.String(http.StatusServiceUnavailable, "Service unavailable")
// 		}

// 	}
// }

func main() {

	fmt.Println("app started successfully ...")

	ctxBackground := context.Background()

	infrastructure.InitDB(ctxBackground, "dsn write in here")

	e := echo.New()

	apiV1Grope := e.Group("/api/v1")

	clients := apiV1Grope.Group("/clients")
	clients.POST("/register", userinterface.RegisterClient())
	clients.POST("/login", userinterface.LoginClient())

	votes := apiV1Grope.Group("/votes")
	votes.GET("/all-votes", userinterface.GetAllVotes())
	votes.POST("/register-vote", userinterface.RegisterNewVote())

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))

}
