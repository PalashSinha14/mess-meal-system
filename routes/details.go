package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func details(c *gin.Context) {
	userId := c.GetInt64("userId")
	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
	})
}
