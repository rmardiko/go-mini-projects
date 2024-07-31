package handler

import (
	"net/http"
	"rmardiko/auth-api/models"
	"rmardiko/auth-api/service"
	"rmardiko/auth-api/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserService service.UserService
	//AuthService service.AuthService
}

func NewAuthHandler(us service.UserService) AuthHandler {

	return AuthHandler{
		UserService: us,
	}
}

func (h *AuthHandler) Hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}

func (h *AuthHandler) Login(c *gin.Context) {

	var user models.User

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// validate username and password
	err := h.UserService.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Credentials"})
		return
	}

	tokenString, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	//c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *AuthHandler) AddUser(c *gin.Context) {

	var newUser models.User

	// Call BindJSON to bind the received JSON to newUser
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	err := h.UserService.AddUser(newUser)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *AuthHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, users)
}
