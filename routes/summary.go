package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/PalashSinha14/mess-meal-system/db"
)

func summaryMeal(c *gin.Context) {
	mealId := c.Param("id")

	row := db.DB.QueryRow(`
	SELECT COUNT(*) FROM meal_confirmations
	WHERE meal_id = ? AND status = 'YES'`, mealId)

	var count int
	row.Scan(&count)

	c.JSON(http.StatusOK, gin.H{
		"meal_id":        mealId,
		"total_confirmed": count,
	})
}
