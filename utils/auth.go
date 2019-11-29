package utils

import (
	"github.com/gin-gonic/gin"
	"moss-service/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	"fmt"
)

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"exp":      time.Now().Add(time.Minute * 60).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tkStr := c.Request.Header.Get("authorization")
		if tkStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		} else {
			token, _ := jwt.Parse(tkStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "Unauthorized",
					})
					c.Abort()
					return nil,fmt.Errorf("not authorization")
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
			} else {
				c.Next()
				return
			}
		}

	}
}
