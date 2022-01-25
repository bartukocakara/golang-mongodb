package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
}