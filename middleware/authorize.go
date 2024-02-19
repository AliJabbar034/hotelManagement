package middleware

import (
	"errors"
	"net/http"

	"github.com/alijabbar034/hotelManagement/types"
	"github.com/alijabbar034/hotelManagement/utils"
	"github.com/gin-gonic/gin"
)

func Authorize(role string) gin.HandlerFunc {

	return func(c *gin.Context) {

		usr, _ := c.Get("user")
		user := usr.(types.User)
		if user.Role != role {
			utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("Unauthorize role"))
			c.Abort()
			return
		}
		c.Next()

	}
}
