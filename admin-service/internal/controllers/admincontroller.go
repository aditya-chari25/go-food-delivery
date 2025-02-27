package controllers

import(
	"admin-service/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
	Role     string
}

func GetUsersData (c *gin.Context){
	JsonData, err := services.GetUsersData()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Users Data": JsonData})
}