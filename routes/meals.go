package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/PalashSinha14/mess-meal-system/models"
)

func getMeals(c *gin.Context) {
	meals,err:=models.GetAllMeals()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meals. Try again later."})
		return
	}
	c.JSON(http.StatusOK, meals)
}

func getMeal(c *gin.Context){
	mealId, err:=strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse meal id."})
		return
	}
	meal, err:=models.GetMealByID(mealId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch meal."})
		return
	}
	c.JSON(http.StatusOK,meal)
}

func createMeal(c *gin.Context){

	var meal models.Meal
	err := c.ShouldBindJSON(&meal)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data"})
		return
	}

	userId := c.GetInt64("userId")
	meal.UserID=userId

	err=meal.Save()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create meal. Try again later."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"Meal Created!", "meal":meal})
}