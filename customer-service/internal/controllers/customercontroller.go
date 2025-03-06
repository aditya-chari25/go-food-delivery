package controllers

import (
	"customer-service/internal/model"
	"customer-service/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context){
	var customerOrder model.Orders
	if err := c.ShouldBindJSON(&customerOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	insertedOrder,err := services.PlaceOrder(customerOrder); 
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"inserted":insertedOrder})
}