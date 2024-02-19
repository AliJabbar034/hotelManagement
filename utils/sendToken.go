package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func SendToken(c *gin.Context, id string) {

	godotenv.Load()
	secret_key := os.Getenv("SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	claims["_id"] = id
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("token", tokenString, int(time.Hour.Seconds()*24*30), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"id":    id,
	})

}
