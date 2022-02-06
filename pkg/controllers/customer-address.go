package controllers

import (
	"restaurant-management/pkg/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var customerAddressCollection *mongo.Collection = database.OpenCollection(database.Client, "customer_addresses")
// var validate = validator.New()

func GetCustomerAddresses() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func GetCustomerAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func CreateCustomerAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func UpdateCustomerAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func DeleteCustomerAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}
