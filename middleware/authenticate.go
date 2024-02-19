package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	AuthStore dbstores.Auth_store
}

func NewAuth(coll dbstores.Auth_store) *Auth {
	return &Auth{
		AuthStore: coll,
	}
}
func (a *Auth) Authenticate(c *gin.Context) {
	secret_key := os.Getenv("SECRET_KEY")
	token, err := c.Cookie("token")
	if err != nil {
		utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("Unauthorize"))
		c.Abort()
		return
	}

	parsed, er := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret_key), nil
	})
	if er != nil {
		utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("Unauthorize"))
		c.Abort()
		return
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || !parsed.Valid {
		utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("Unauthorize"))
		c.Abort()
		return
	}
	id := claims["_id"].(string)
	user, err := a.AuthStore.FindUser(id)
	fmt.Println(user, err)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		c.Abort()
		return

	}
	c.Set("user", *user)
	c.Next()

}
