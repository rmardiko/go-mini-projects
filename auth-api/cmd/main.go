package main

import (
	"log"
	"rmardiko/auth-api/handler"
	"rmardiko/auth-api/middleware"
	"rmardiko/auth-api/models"
	"rmardiko/auth-api/repository"
	"rmardiko/auth-api/service"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var users = []models.User{
	{ID: "1", Username: "andrew", Password: "easypassword"},
	{ID: "2", Username: "paul", Password: "plainpassword"},
	{ID: "3", Username: "jonathan", Password: "abcpassword"},
}

func main() {

	userRepo := repository.NewSliceUserRepository()
	userService := service.NewUserService(userRepo)

	// add users
	for _, u := range users {
		err := userService.AddUser(u)
		if err != nil {
			log.Println(err)
		}
	}

	handler := handler.NewAuthHandler(userService)

	router := gin.Default()

	// dummy endpoint
	router.GET("/hello", handler.Hello)

	// auth endpoints
	router.POST("/login", handler.Login)

	// users CRUD endpoints
	router.GET("/users", middleware.AuthenticateToken(), handler.GetUsers)
	router.POST("/users", middleware.AuthenticateToken(), handler.AddUser)

	router.Run("localhost:8080")
}
