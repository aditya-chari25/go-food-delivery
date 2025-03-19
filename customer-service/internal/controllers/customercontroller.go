package controllers

import (
	"customer-service/internal/model"
	"customer-service/internal/services"
	"customer-service/internal/validators"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertOrder(c *gin.Context){
	var customerOrder model.Orders
	if err := c.ShouldBindJSON(&customerOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.ValidateOrder(customerOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// log.Printf("order in controller %v", customerOrder)
	insertedOrder,err := services.PlaceOrder(customerOrder); 
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"inserted":insertedOrder})
}

func ReturnMenu(c *gin.Context){
	var restMenu model.RestMenu
	if err := c.ShouldBindJSON(&restMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	returnmenu,err := services.RestMenu(restMenu)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})	
	}
	
	c.JSON(http.StatusOK, gin.H{"Menu":returnmenu})
}

