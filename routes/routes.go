package routes 

import (
	"github.com/gin-gonic/gin"
	"github.com/PalashSinha14/mess-meal-system/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/meals", getMeals) //to get list of all available meals
	server.GET("/meal/:id", getMeal) //to see any1 particular meal

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/mydetails", details) //to get user details who has logged in
	authenticated.POST("/meals", createMeal) //endpoint for staff to add meal to menu for students
	authenticated.GET("/meal/confirm", getConfirmMeal) //endpoint to get status of meal
	authenticated.POST("/meal/confirm", confirmMeal) //endpoint to post status of meal
	authenticated.GET("/staff/meal/:id/summary", summaryMeal) //endpoint for staff to get summary of how much food to make today
	
}