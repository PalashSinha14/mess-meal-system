package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/PalashSinha14/mess-meal-system/models"
)

func confirmMeal(c *gin.Context) {
	userId := c.GetInt64("userId")

	var input struct {
		MealID int64  `json:"meal_id"`
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	confirmation := models.MealConfirmation{
		UserID: userId,
		MealID: input.MealID,
		Status: input.Status,
	}

	err := confirmation.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not confirm meal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meal confirmation saved"})
}

func getConfirmMeal(c *gin.Context) {
	userId := c.GetInt64("userId")
	data, err := models.GetUserConfirmation(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch confirmations"})
		return
	}
	c.JSON(http.StatusOK, data)
}
