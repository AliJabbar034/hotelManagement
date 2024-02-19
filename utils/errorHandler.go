package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, statusCode int, err error) {

	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
	return
}
