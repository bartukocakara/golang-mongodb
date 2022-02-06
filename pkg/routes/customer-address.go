package routes

import (
	"restaurant-management/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerAddressRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/customer-addresses", controllers.GetCustomerAddresses()) //++
	incomingRoutes.POST("/customer-address", controllers.CreateCustomerAddress()) //++
	incomingRoutes.GET("/customer-address/:food_id", controllers.GetCustomerAddress())
	incomingRoutes.PATCH("/customer-address/:food_id", controllers.UpdateCustomerAddress())
	incomingRoutes.DELETE("/customer-address/:food_id", controllers.DeleteCustomerAddress())
}