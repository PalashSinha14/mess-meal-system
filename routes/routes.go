package routes 

import (
	"github.com/gin-gonic/gin"
	"github.com/PalashSinha14/mess-meal-system/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	//server.GET("/meals", getMeals)
	//server.GET("/meal/:id", getMeal)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	//authenticated.POST("/meals", createMeal)
	server.POST("/signup", signup)
	server.POST("/login", login)
}