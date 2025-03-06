package controllers

import (
	"auth-service/internal/services"
	"auth-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handler
func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Verfication(c *gin.Context){
	var credentials struct{
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	retusername,err := services.VerifyUser(credentials.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": retusername})
}

func Signup(c *gin.Context){
	var customer model.SignUp
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inscustomer,err := services.SignUser(customer); 
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"inserted":inscustomer})
}

