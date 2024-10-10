package handlers 

import (
	// "github.com/code-guy21/PingUp/server/models"
	// "github.com/code-guy21/PingUp/server/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context){
	// var user models.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	// 	return
	// }

	// err := auth.RegisterUser(&user)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "User registered sucesssfully"})
}

func LoginUser(c *gin.Context){
	// var loginData struct {
	// 	Email string `json:"email"` 
	// 	Password string `json:"password"`
	// }

	// if err := c.ShouldBindJSON(&loginData); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 	return
	// }

	// token, err := auth.AuthenticateUser(loginData.Email, loginData.Password)

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }
	
	c.JSON(http.StatusOK, gin.H{"message":"token"/*: token*/})
}